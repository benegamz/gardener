// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/gardener/gardener/pkg/client/settings/clientset/versioned/typed/settings/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSettingsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSettingsV1alpha1) ClusterOpenIDConnectPresets() v1alpha1.ClusterOpenIDConnectPresetInterface {
	return &FakeClusterOpenIDConnectPresets{c}
}

func (c *FakeSettingsV1alpha1) OpenIDConnectPresets(namespace string) v1alpha1.OpenIDConnectPresetInterface {
	return &FakeOpenIDConnectPresets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSettingsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
