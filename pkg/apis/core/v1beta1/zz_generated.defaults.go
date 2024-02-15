//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by defaulter-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&CloudProfile{}, func(obj interface{}) { SetObjectDefaults_CloudProfile(obj.(*CloudProfile)) })
	scheme.AddTypeDefaultingFunc(&CloudProfileList{}, func(obj interface{}) { SetObjectDefaults_CloudProfileList(obj.(*CloudProfileList)) })
	scheme.AddTypeDefaultingFunc(&ControllerRegistration{}, func(obj interface{}) { SetObjectDefaults_ControllerRegistration(obj.(*ControllerRegistration)) })
	scheme.AddTypeDefaultingFunc(&ControllerRegistrationList{}, func(obj interface{}) { SetObjectDefaults_ControllerRegistrationList(obj.(*ControllerRegistrationList)) })
	scheme.AddTypeDefaultingFunc(&PrivateCloudProfile{}, func(obj interface{}) { SetObjectDefaults_PrivateCloudProfile(obj.(*PrivateCloudProfile)) })
	scheme.AddTypeDefaultingFunc(&PrivateCloudProfileList{}, func(obj interface{}) { SetObjectDefaults_PrivateCloudProfileList(obj.(*PrivateCloudProfileList)) })
	scheme.AddTypeDefaultingFunc(&Project{}, func(obj interface{}) { SetObjectDefaults_Project(obj.(*Project)) })
	scheme.AddTypeDefaultingFunc(&ProjectList{}, func(obj interface{}) { SetObjectDefaults_ProjectList(obj.(*ProjectList)) })
	scheme.AddTypeDefaultingFunc(&SecretBinding{}, func(obj interface{}) { SetObjectDefaults_SecretBinding(obj.(*SecretBinding)) })
	scheme.AddTypeDefaultingFunc(&SecretBindingList{}, func(obj interface{}) { SetObjectDefaults_SecretBindingList(obj.(*SecretBindingList)) })
	scheme.AddTypeDefaultingFunc(&Seed{}, func(obj interface{}) { SetObjectDefaults_Seed(obj.(*Seed)) })
	scheme.AddTypeDefaultingFunc(&SeedList{}, func(obj interface{}) { SetObjectDefaults_SeedList(obj.(*SeedList)) })
	scheme.AddTypeDefaultingFunc(&Shoot{}, func(obj interface{}) { SetObjectDefaults_Shoot(obj.(*Shoot)) })
	scheme.AddTypeDefaultingFunc(&ShootList{}, func(obj interface{}) { SetObjectDefaults_ShootList(obj.(*ShootList)) })
	return nil
}

func SetObjectDefaults_CloudProfile(in *CloudProfile) {
	for i := range in.Spec.MachineImages {
		a := &in.Spec.MachineImages[i]
		SetDefaults_MachineImage(a)
		for j := range a.Versions {
			b := &a.Versions[j]
			SetDefaults_MachineImageVersion(b)
		}
	}
	for i := range in.Spec.MachineTypes {
		a := &in.Spec.MachineTypes[i]
		SetDefaults_MachineType(a)
	}
	for i := range in.Spec.VolumeTypes {
		a := &in.Spec.VolumeTypes[i]
		SetDefaults_VolumeType(a)
	}
}

func SetObjectDefaults_CloudProfileList(in *CloudProfileList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_CloudProfile(a)
	}
}

func SetObjectDefaults_ControllerRegistration(in *ControllerRegistration) {
	for i := range in.Spec.Resources {
		a := &in.Spec.Resources[i]
		SetDefaults_ControllerResource(a)
		if a.Lifecycle != nil {
			SetDefaults_ControllerResourceLifecycle(a.Lifecycle)
		}
	}
	if in.Spec.Deployment != nil {
		SetDefaults_ControllerRegistrationDeployment(in.Spec.Deployment)
	}
}

func SetObjectDefaults_ControllerRegistrationList(in *ControllerRegistrationList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ControllerRegistration(a)
	}
}

func SetObjectDefaults_PrivateCloudProfile(in *PrivateCloudProfile) {
	for i := range in.Spec.MachineImages {
		a := &in.Spec.MachineImages[i]
		SetDefaults_MachineImage(a)
		for j := range a.Versions {
			b := &a.Versions[j]
			SetDefaults_MachineImageVersion(b)
		}
	}
	for i := range in.Spec.MachineTypes {
		a := &in.Spec.MachineTypes[i]
		SetDefaults_MachineType(a)
	}
	for i := range in.Spec.VolumeTypes {
		a := &in.Spec.VolumeTypes[i]
		SetDefaults_VolumeType(a)
	}
}

func SetObjectDefaults_PrivateCloudProfileList(in *PrivateCloudProfileList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_PrivateCloudProfile(a)
	}
}

func SetObjectDefaults_Project(in *Project) {
	SetDefaults_Project(in)
	for i := range in.Spec.Members {
		a := &in.Spec.Members[i]
		SetDefaults_ProjectMember(a)
	}
}

func SetObjectDefaults_ProjectList(in *ProjectList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Project(a)
	}
}

func SetObjectDefaults_SecretBinding(in *SecretBinding) {
	SetDefaults_SecretBinding(in)
}

func SetObjectDefaults_SecretBindingList(in *SecretBindingList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_SecretBinding(a)
	}
}

func SetObjectDefaults_Seed(in *Seed) {
	SetDefaults_Seed(in)
	SetDefaults_SeedNetworks(&in.Spec.Networks)
	if in.Spec.Settings != nil {
		SetDefaults_SeedSettings(in.Spec.Settings)
		if in.Spec.Settings.DependencyWatchdog != nil {
			SetDefaults_SeedSettingDependencyWatchdog(in.Spec.Settings.DependencyWatchdog)
		}
	}
}

func SetObjectDefaults_SeedList(in *SeedList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Seed(a)
	}
}

func SetObjectDefaults_Shoot(in *Shoot) {
	SetDefaults_Shoot(in)
	if in.Spec.Addons != nil {
		if in.Spec.Addons.NginxIngress != nil {
			SetDefaults_NginxIngress(in.Spec.Addons.NginxIngress)
		}
	}
	if in.Spec.Kubernetes.ClusterAutoscaler != nil {
		SetDefaults_ClusterAutoscaler(in.Spec.Kubernetes.ClusterAutoscaler)
	}
	if in.Spec.Kubernetes.KubeAPIServer != nil {
		SetDefaults_KubeAPIServerConfig(in.Spec.Kubernetes.KubeAPIServer)
	}
	if in.Spec.Kubernetes.VerticalPodAutoscaler != nil {
		SetDefaults_VerticalPodAutoscaler(in.Spec.Kubernetes.VerticalPodAutoscaler)
	}
	if in.Spec.Networking != nil {
		SetDefaults_Networking(in.Spec.Networking)
	}
	if in.Spec.Maintenance != nil {
		SetDefaults_Maintenance(in.Spec.Maintenance)
	}
	for i := range in.Spec.Provider.Workers {
		a := &in.Spec.Provider.Workers[i]
		SetDefaults_Worker(a)
	}
}

func SetObjectDefaults_ShootList(in *ShootList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Shoot(a)
	}
}
