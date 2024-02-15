// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/gardener/gardener/pkg/client/settings/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterOpenIDConnectPresets returns a ClusterOpenIDConnectPresetInformer.
	ClusterOpenIDConnectPresets() ClusterOpenIDConnectPresetInformer
	// OpenIDConnectPresets returns a OpenIDConnectPresetInformer.
	OpenIDConnectPresets() OpenIDConnectPresetInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterOpenIDConnectPresets returns a ClusterOpenIDConnectPresetInformer.
func (v *version) ClusterOpenIDConnectPresets() ClusterOpenIDConnectPresetInformer {
	return &clusterOpenIDConnectPresetInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// OpenIDConnectPresets returns a OpenIDConnectPresetInformer.
func (v *version) OpenIDConnectPresets() OpenIDConnectPresetInformer {
	return &openIDConnectPresetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
