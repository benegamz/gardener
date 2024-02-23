// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	"fmt"

	core "github.com/gardener/gardener/pkg/apis/core"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=core.gardener.cloud, Version=internalVersion
	case core.SchemeGroupVersion.WithResource("backupbuckets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().BackupBuckets().Informer()}, nil
	case core.SchemeGroupVersion.WithResource("backupentries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().BackupEntries().Informer()}, nil
	case core.SchemeGroupVersion.WithResource("cloudprofiles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().CloudProfiles().Informer()}, nil
	case core.SchemeGroupVersion.WithResource("controllerdeployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().ControllerDeployments().Informer()}, nil
	case core.SchemeGroupVersion.WithResource("controllerinstallations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().ControllerInstallations().Informer()}, nil
	case core.SchemeGroupVersion.WithResource("controllerregistrations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().ControllerRegistrations().Informer()}, nil
	case core.SchemeGroupVersion.WithResource("exposureclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().ExposureClasses().Informer()}, nil
	case core.SchemeGroupVersion.WithResource("internalsecrets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().InternalSecrets().Informer()}, nil
	case core.SchemeGroupVersion.WithResource("namespacedcloudprofiles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().NamespacedCloudProfiles().Informer()}, nil
	case core.SchemeGroupVersion.WithResource("projects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Projects().Informer()}, nil
	case core.SchemeGroupVersion.WithResource("quotas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Quotas().Informer()}, nil
	case core.SchemeGroupVersion.WithResource("secretbindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().SecretBindings().Informer()}, nil
	case core.SchemeGroupVersion.WithResource("seeds"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Seeds().Informer()}, nil
	case core.SchemeGroupVersion.WithResource("shoots"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Shoots().Informer()}, nil
	case core.SchemeGroupVersion.WithResource("shootstates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().ShootStates().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
