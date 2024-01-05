/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/myoperator/multiclusteroperator/pkg/apis/multicluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMultiClusters implements MultiClusterInterface
type FakeMultiClusters struct {
	Fake *FakeMulitclusterV1alpha1
	ns   string
}

var multiclustersResource = schema.GroupVersionResource{Group: "mulitcluster.practice.com", Version: "v1alpha1", Resource: "multiclusters"}

var multiclustersKind = schema.GroupVersionKind{Group: "mulitcluster.practice.com", Version: "v1alpha1", Kind: "MultiCluster"}

// Get takes name of the multiCluster, and returns the corresponding multiCluster object, and an error if there is any.
func (c *FakeMultiClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MultiCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(multiclustersResource, c.ns, name), &v1alpha1.MultiCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiCluster), err
}

// List takes label and field selectors, and returns the list of MultiClusters that match those selectors.
func (c *FakeMultiClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MultiClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(multiclustersResource, multiclustersKind, c.ns, opts), &v1alpha1.MultiClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MultiClusterList{ListMeta: obj.(*v1alpha1.MultiClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.MultiClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested multiClusters.
func (c *FakeMultiClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(multiclustersResource, c.ns, opts))

}

// Create takes the representation of a multiCluster and creates it.  Returns the server's representation of the multiCluster, and an error, if there is any.
func (c *FakeMultiClusters) Create(ctx context.Context, multiCluster *v1alpha1.MultiCluster, opts v1.CreateOptions) (result *v1alpha1.MultiCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(multiclustersResource, c.ns, multiCluster), &v1alpha1.MultiCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiCluster), err
}

// Update takes the representation of a multiCluster and updates it. Returns the server's representation of the multiCluster, and an error, if there is any.
func (c *FakeMultiClusters) Update(ctx context.Context, multiCluster *v1alpha1.MultiCluster, opts v1.UpdateOptions) (result *v1alpha1.MultiCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(multiclustersResource, c.ns, multiCluster), &v1alpha1.MultiCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMultiClusters) UpdateStatus(ctx context.Context, multiCluster *v1alpha1.MultiCluster, opts v1.UpdateOptions) (*v1alpha1.MultiCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(multiclustersResource, "status", c.ns, multiCluster), &v1alpha1.MultiCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiCluster), err
}

// Delete takes name of the multiCluster and deletes it. Returns an error if one occurs.
func (c *FakeMultiClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(multiclustersResource, c.ns, name), &v1alpha1.MultiCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMultiClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(multiclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MultiClusterList{})
	return err
}

// Patch applies the patch and returns the patched multiCluster.
func (c *FakeMultiClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MultiCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(multiclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.MultiCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiCluster), err
}
