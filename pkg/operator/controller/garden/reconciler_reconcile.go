// Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package garden

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	operatorv1alpha1 "github.com/gardener/gardener/pkg/apis/operator/v1alpha1"
	resourcesv1alpha1 "github.com/gardener/gardener/pkg/apis/resources/v1alpha1"
	"github.com/gardener/gardener/pkg/client/kubernetes"
	"github.com/gardener/gardener/pkg/controllerutils"
	"github.com/gardener/gardener/pkg/operation/botanist/component"
	"github.com/gardener/gardener/pkg/operation/botanist/component/hvpa"
	"github.com/gardener/gardener/pkg/operation/botanist/component/vpa"
	"github.com/gardener/gardener/pkg/utils/flow"
	secretutils "github.com/gardener/gardener/pkg/utils/secrets"
	secretsmanager "github.com/gardener/gardener/pkg/utils/secrets/manager"
)

func (r *Reconciler) reconcile(
	ctx context.Context,
	log logr.Logger,
	garden *operatorv1alpha1.Garden,
	secretsManager secretsmanager.Interface,
) (
	reconcile.Result,
	error,
) {
	applier := kubernetes.NewApplier(r.RuntimeClient, r.RuntimeClient.RESTMapper())

	if !controllerutil.ContainsFinalizer(garden, finalizerName) {
		log.Info("Adding finalizer")
		if err := controllerutils.AddFinalizers(ctx, r.RuntimeClient, garden, finalizerName); err != nil {
			return reconcile.Result{}, err
		}
	}

	// VPA is a prerequisite. If it's enabled then we deploy the CRD (and later also the related components) as part of
	// the flow. However, when it's disabled then we check whether it is indeed available (and fail, otherwise).
	if !vpaEnabled(garden.Spec.RuntimeCluster.Settings) {
		if _, err := r.RuntimeClient.RESTMapper().RESTMapping(schema.GroupKind{Group: "autoscaling.k8s.io", Kind: "VerticalPodAutoscaler"}); err != nil {
			return reconcile.Result{}, fmt.Errorf("VPA is required for runtime cluster but CRD is not installed: %s", err)
		}
	}

	// create + label namespace
	namespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: r.GardenNamespace}}
	log.Info("Labeling and annotating namespace", "namespaceName", namespace.Name)
	if _, err := controllerutils.CreateOrGetAndMergePatch(ctx, r.RuntimeClient, namespace, func() error {
		metav1.SetMetaDataLabel(&namespace.ObjectMeta, resourcesv1alpha1.HighAvailabilityConfigConsider, "true")
		metav1.SetMetaDataAnnotation(&namespace.ObjectMeta, resourcesv1alpha1.HighAvailabilityConfigZones, strings.Join(garden.Spec.RuntimeCluster.Provider.Zones, ","))
		return nil
	}); err != nil {
		return reconcile.Result{}, err
	}

	log.Info("Generating general CA certificate for runtime cluster")
	if _, err := secretsManager.Generate(ctx, &secretutils.CertificateSecretConfig{
		Name:       operatorv1alpha1.SecretNameCARuntime,
		CommonName: "garden-runtime",
		CertType:   secretutils.CACert,
		Validity:   pointer.Duration(30 * 24 * time.Hour),
	}, secretsmanager.Rotate(secretsmanager.KeepOld), secretsmanager.IgnoreOldSecretsAfter(24*time.Hour)); err != nil {
		return reconcile.Result{}, err
	}

	log.Info("Instantiating component deployers")
	vpaCRD := vpa.NewCRD(applier, nil)
	hvpaCRD := hvpa.NewCRD(applier)
	gardenerResourceManager, err := r.newGardenerResourceManager(garden, secretsManager)
	if err != nil {
		return reconcile.Result{}, err
	}
	system := r.newSystem()
	verticalPodAutoscaler, err := r.newVerticalPodAutoscaler(garden, secretsManager)
	if err != nil {
		return reconcile.Result{}, err
	}
	hvpaController, err := r.newHVPA()
	if err != nil {
		return reconcile.Result{}, err
	}
	etcdDruid, err := r.newEtcdDruid()
	if err != nil {
		return reconcile.Result{}, err
	}

	if !hvpaEnabled() {
		hvpaCRD = component.OpDestroy(hvpaCRD)
	}

	var (
		g            = flow.NewGraph("Garden reconciliation")
		deployVPACRD = g.Add(flow.Task{
			Name: "Deploying custom resource definition for VPA",
			Fn:   flow.TaskFn(vpaCRD.Deploy).DoIf(vpaEnabled(garden.Spec.RuntimeCluster.Settings)),
		})
		reconcileHVPACRD = g.Add(flow.Task{
			Name: "Reconciling custom resource definition for HVPA",
			Fn:   hvpaCRD.Deploy,
		})
		deployGardenerResourceManager = g.Add(flow.Task{
			Name:         "Deploying and waiting for gardener-resource-manager to be healthy",
			Fn:           component.OpWait(gardenerResourceManager).Deploy,
			Dependencies: flow.NewTaskIDs(deployVPACRD, reconcileHVPACRD),
		})
		_ = g.Add(flow.Task{
			Name:         "Deploying system resources",
			Fn:           system.Deploy,
			Dependencies: flow.NewTaskIDs(deployGardenerResourceManager),
		})
		_ = g.Add(flow.Task{
			Name:         "Deploying Kubernetes vertical pod autoscaler",
			Fn:           verticalPodAutoscaler.Deploy,
			Dependencies: flow.NewTaskIDs(deployGardenerResourceManager),
		})
		_ = g.Add(flow.Task{
			Name:         "Deploying HVPA controller",
			Fn:           hvpaController.Deploy,
			Dependencies: flow.NewTaskIDs(deployGardenerResourceManager),
		})
		_ = g.Add(flow.Task{
			Name:         "Deploying ETCD Druid",
			Fn:           etcdDruid.Deploy,
			Dependencies: flow.NewTaskIDs(deployGardenerResourceManager),
		})
	)

	if err := g.Compile().Run(ctx, flow.Opts{Log: log}); err != nil {
		return reconcile.Result{}, flow.Errors(err)
	}

	return reconcile.Result{}, secretsManager.Cleanup(ctx)
}