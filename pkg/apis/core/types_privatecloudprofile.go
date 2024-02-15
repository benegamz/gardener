// Copyright 2024 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package core

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PrivateCloudProfile represents certain properties about a provider environment.
type PrivateCloudProfile struct {
	metav1.TypeMeta
	// Standard object metadata.
	// +optional
	metav1.ObjectMeta
	// Spec defines the provider environment properties.
	// +optional
	Spec PrivateCloudProfileSpec
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PrivateCloudProfileList is a collection of CloudProfiles.
type PrivateCloudProfileList struct {
	metav1.TypeMeta
	// Standard list object metadata.
	// +optional
	metav1.ListMeta
	// Items is the list of CloudProfiles.
	Items []PrivateCloudProfile
}

// PrivateCloudProfileSpec is the specification of a PrivateCloudProfile.
// It must contain exactly one of its defined keys.
type PrivateCloudProfileSpec struct {
	// CABundle is a certificate bundle which will be installed onto every host machine of shoot cluster targeting this profile.
	// +optional
	CABundle *string
	// Kubernetes contains constraints regarding allowed values of the 'kubernetes' block in the Shoot specification.
	// +optional
	Kubernetes *KubernetesSettings
	// MachineImages contains constraints regarding allowed values for machine images in the Shoot specification.
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +optional
	MachineImages []MachineImage
	// MachineTypes contains constraints regarding allowed values for machine types in the 'workers' block in the Shoot specification.
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +optional
	MachineTypes []MachineType
	// Regions contains constraints regarding allowed values for regions and zones.
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +optional
	Regions []Region
	// SeedSelector contains an optional list of labels on `Seed` resources that marks those seeds whose shoots may use this provider profile.
	// An empty list means that all seeds of the same provider type are supported.
	// This is useful for environments that are of the same type (like openstack) but may have different "instances"/landscapes.
	// Optionally a list of possible providers can be added to enable cross-provider scheduling. By default, the provider
	// type of the seed must match the shoot's provider.
	// +optional
	SeedSelector *SeedSelector
	// VolumeTypes contains constraints regarding allowed values for volume types in the 'workers' block in the Shoot specification.
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +optional
	VolumeTypes []VolumeType
	// A pointer to the PrivateCloudProfiles parent CloudProfile
	// +optional
	Parent string
}
