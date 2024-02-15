// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gardener/gardener/pkg/apis/settings/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterOpenIDConnectPresets implements ClusterOpenIDConnectPresetInterface
type FakeClusterOpenIDConnectPresets struct {
	Fake *FakeSettingsV1alpha1
}

var clusteropenidconnectpresetsResource = v1alpha1.SchemeGroupVersion.WithResource("clusteropenidconnectpresets")

var clusteropenidconnectpresetsKind = v1alpha1.SchemeGroupVersion.WithKind("ClusterOpenIDConnectPreset")

// Get takes name of the clusterOpenIDConnectPreset, and returns the corresponding clusterOpenIDConnectPreset object, and an error if there is any.
func (c *FakeClusterOpenIDConnectPresets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterOpenIDConnectPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusteropenidconnectpresetsResource, name), &v1alpha1.ClusterOpenIDConnectPreset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterOpenIDConnectPreset), err
}

// List takes label and field selectors, and returns the list of ClusterOpenIDConnectPresets that match those selectors.
func (c *FakeClusterOpenIDConnectPresets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterOpenIDConnectPresetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusteropenidconnectpresetsResource, clusteropenidconnectpresetsKind, opts), &v1alpha1.ClusterOpenIDConnectPresetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterOpenIDConnectPresetList{ListMeta: obj.(*v1alpha1.ClusterOpenIDConnectPresetList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterOpenIDConnectPresetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterOpenIDConnectPresets.
func (c *FakeClusterOpenIDConnectPresets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusteropenidconnectpresetsResource, opts))
}

// Create takes the representation of a clusterOpenIDConnectPreset and creates it.  Returns the server's representation of the clusterOpenIDConnectPreset, and an error, if there is any.
func (c *FakeClusterOpenIDConnectPresets) Create(ctx context.Context, clusterOpenIDConnectPreset *v1alpha1.ClusterOpenIDConnectPreset, opts v1.CreateOptions) (result *v1alpha1.ClusterOpenIDConnectPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusteropenidconnectpresetsResource, clusterOpenIDConnectPreset), &v1alpha1.ClusterOpenIDConnectPreset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterOpenIDConnectPreset), err
}

// Update takes the representation of a clusterOpenIDConnectPreset and updates it. Returns the server's representation of the clusterOpenIDConnectPreset, and an error, if there is any.
func (c *FakeClusterOpenIDConnectPresets) Update(ctx context.Context, clusterOpenIDConnectPreset *v1alpha1.ClusterOpenIDConnectPreset, opts v1.UpdateOptions) (result *v1alpha1.ClusterOpenIDConnectPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusteropenidconnectpresetsResource, clusterOpenIDConnectPreset), &v1alpha1.ClusterOpenIDConnectPreset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterOpenIDConnectPreset), err
}

// Delete takes name of the clusterOpenIDConnectPreset and deletes it. Returns an error if one occurs.
func (c *FakeClusterOpenIDConnectPresets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusteropenidconnectpresetsResource, name, opts), &v1alpha1.ClusterOpenIDConnectPreset{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterOpenIDConnectPresets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusteropenidconnectpresetsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterOpenIDConnectPresetList{})
	return err
}

// Patch applies the patch and returns the patched clusterOpenIDConnectPreset.
func (c *FakeClusterOpenIDConnectPresets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterOpenIDConnectPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusteropenidconnectpresetsResource, name, pt, data, subresources...), &v1alpha1.ClusterOpenIDConnectPreset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterOpenIDConnectPreset), err
}
