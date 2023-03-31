// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kubewharf/kubeadmiral/pkg/apis/core/v1alpha1"
	scheme "github.com/kubewharf/kubeadmiral/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterPropagatedVersionsGetter has a method to return a ClusterPropagatedVersionInterface.
// A group's client should implement this interface.
type ClusterPropagatedVersionsGetter interface {
	ClusterPropagatedVersions() ClusterPropagatedVersionInterface
}

// ClusterPropagatedVersionInterface has methods to work with ClusterPropagatedVersion resources.
type ClusterPropagatedVersionInterface interface {
	Create(ctx context.Context, clusterPropagatedVersion *v1alpha1.ClusterPropagatedVersion, opts v1.CreateOptions) (*v1alpha1.ClusterPropagatedVersion, error)
	Update(ctx context.Context, clusterPropagatedVersion *v1alpha1.ClusterPropagatedVersion, opts v1.UpdateOptions) (*v1alpha1.ClusterPropagatedVersion, error)
	UpdateStatus(ctx context.Context, clusterPropagatedVersion *v1alpha1.ClusterPropagatedVersion, opts v1.UpdateOptions) (*v1alpha1.ClusterPropagatedVersion, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterPropagatedVersion, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterPropagatedVersionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterPropagatedVersion, err error)
	ClusterPropagatedVersionExpansion
}

// clusterPropagatedVersions implements ClusterPropagatedVersionInterface
type clusterPropagatedVersions struct {
	client rest.Interface
}

// newClusterPropagatedVersions returns a ClusterPropagatedVersions
func newClusterPropagatedVersions(c *CoreV1alpha1Client) *clusterPropagatedVersions {
	return &clusterPropagatedVersions{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterPropagatedVersion, and returns the corresponding clusterPropagatedVersion object, and an error if there is any.
func (c *clusterPropagatedVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterPropagatedVersion, err error) {
	result = &v1alpha1.ClusterPropagatedVersion{}
	err = c.client.Get().
		Resource("clusterpropagatedversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterPropagatedVersions that match those selectors.
func (c *clusterPropagatedVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterPropagatedVersionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterPropagatedVersionList{}
	err = c.client.Get().
		Resource("clusterpropagatedversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterPropagatedVersions.
func (c *clusterPropagatedVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterpropagatedversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterPropagatedVersion and creates it.  Returns the server's representation of the clusterPropagatedVersion, and an error, if there is any.
func (c *clusterPropagatedVersions) Create(ctx context.Context, clusterPropagatedVersion *v1alpha1.ClusterPropagatedVersion, opts v1.CreateOptions) (result *v1alpha1.ClusterPropagatedVersion, err error) {
	result = &v1alpha1.ClusterPropagatedVersion{}
	err = c.client.Post().
		Resource("clusterpropagatedversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterPropagatedVersion).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterPropagatedVersion and updates it. Returns the server's representation of the clusterPropagatedVersion, and an error, if there is any.
func (c *clusterPropagatedVersions) Update(ctx context.Context, clusterPropagatedVersion *v1alpha1.ClusterPropagatedVersion, opts v1.UpdateOptions) (result *v1alpha1.ClusterPropagatedVersion, err error) {
	result = &v1alpha1.ClusterPropagatedVersion{}
	err = c.client.Put().
		Resource("clusterpropagatedversions").
		Name(clusterPropagatedVersion.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterPropagatedVersion).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterPropagatedVersions) UpdateStatus(ctx context.Context, clusterPropagatedVersion *v1alpha1.ClusterPropagatedVersion, opts v1.UpdateOptions) (result *v1alpha1.ClusterPropagatedVersion, err error) {
	result = &v1alpha1.ClusterPropagatedVersion{}
	err = c.client.Put().
		Resource("clusterpropagatedversions").
		Name(clusterPropagatedVersion.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterPropagatedVersion).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterPropagatedVersion and deletes it. Returns an error if one occurs.
func (c *clusterPropagatedVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterpropagatedversions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterPropagatedVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterpropagatedversions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterPropagatedVersion.
func (c *clusterPropagatedVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterPropagatedVersion, err error) {
	result = &v1alpha1.ClusterPropagatedVersion{}
	err = c.client.Patch(pt).
		Resource("clusterpropagatedversions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}