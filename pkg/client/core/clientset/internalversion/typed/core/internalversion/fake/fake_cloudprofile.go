// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	core "github.com/gardener/gardener/pkg/apis/core"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudProfiles implements CloudProfileInterface
type FakeCloudProfiles struct {
	Fake *FakeCore
}

var cloudprofilesResource = core.SchemeGroupVersion.WithResource("cloudprofiles")

var cloudprofilesKind = core.SchemeGroupVersion.WithKind("CloudProfile")

// Get takes name of the cloudProfile, and returns the corresponding cloudProfile object, and an error if there is any.
func (c *FakeCloudProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *core.CloudProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(cloudprofilesResource, name), &core.CloudProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.CloudProfile), err
}

// List takes label and field selectors, and returns the list of CloudProfiles that match those selectors.
func (c *FakeCloudProfiles) List(ctx context.Context, opts v1.ListOptions) (result *core.CloudProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(cloudprofilesResource, cloudprofilesKind, opts), &core.CloudProfileList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &core.CloudProfileList{ListMeta: obj.(*core.CloudProfileList).ListMeta}
	for _, item := range obj.(*core.CloudProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudProfiles.
func (c *FakeCloudProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(cloudprofilesResource, opts))
}

// Create takes the representation of a cloudProfile and creates it.  Returns the server's representation of the cloudProfile, and an error, if there is any.
func (c *FakeCloudProfiles) Create(ctx context.Context, cloudProfile *core.CloudProfile, opts v1.CreateOptions) (result *core.CloudProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(cloudprofilesResource, cloudProfile), &core.CloudProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.CloudProfile), err
}

// Update takes the representation of a cloudProfile and updates it. Returns the server's representation of the cloudProfile, and an error, if there is any.
func (c *FakeCloudProfiles) Update(ctx context.Context, cloudProfile *core.CloudProfile, opts v1.UpdateOptions) (result *core.CloudProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(cloudprofilesResource, cloudProfile), &core.CloudProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.CloudProfile), err
}

// Delete takes name of the cloudProfile and deletes it. Returns an error if one occurs.
func (c *FakeCloudProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(cloudprofilesResource, name, opts), &core.CloudProfile{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(cloudprofilesResource, listOpts)

	_, err := c.Fake.Invokes(action, &core.CloudProfileList{})
	return err
}

// Patch applies the patch and returns the patched cloudProfile.
func (c *FakeCloudProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *core.CloudProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(cloudprofilesResource, name, pt, data, subresources...), &core.CloudProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*core.CloudProfile), err
}
