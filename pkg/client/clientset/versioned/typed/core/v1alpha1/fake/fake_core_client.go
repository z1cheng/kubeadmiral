// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kubewharf/kubeadmiral/pkg/client/clientset/versioned/typed/core/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCoreV1alpha1 struct {
	*testing.Fake
}

func (c *FakeCoreV1alpha1) ClusterOverridePolicies() v1alpha1.ClusterOverridePolicyInterface {
	return &FakeClusterOverridePolicies{c}
}

func (c *FakeCoreV1alpha1) ClusterPropagatedVersions() v1alpha1.ClusterPropagatedVersionInterface {
	return &FakeClusterPropagatedVersions{c}
}

func (c *FakeCoreV1alpha1) ClusterPropagationPolicies() v1alpha1.ClusterPropagationPolicyInterface {
	return &FakeClusterPropagationPolicies{c}
}

func (c *FakeCoreV1alpha1) FederatedClusters() v1alpha1.FederatedClusterInterface {
	return &FakeFederatedClusters{c}
}

func (c *FakeCoreV1alpha1) FederatedTypeConfigs() v1alpha1.FederatedTypeConfigInterface {
	return &FakeFederatedTypeConfigs{c}
}

func (c *FakeCoreV1alpha1) OverridePolicies(namespace string) v1alpha1.OverridePolicyInterface {
	return &FakeOverridePolicies{c, namespace}
}

func (c *FakeCoreV1alpha1) PropagatedVersions(namespace string) v1alpha1.PropagatedVersionInterface {
	return &FakePropagatedVersions{c, namespace}
}

func (c *FakeCoreV1alpha1) PropagationPolicies(namespace string) v1alpha1.PropagationPolicyInterface {
	return &FakePropagationPolicies{c, namespace}
}

func (c *FakeCoreV1alpha1) SchedulingProfiles() v1alpha1.SchedulingProfileInterface {
	return &FakeSchedulingProfiles{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCoreV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}