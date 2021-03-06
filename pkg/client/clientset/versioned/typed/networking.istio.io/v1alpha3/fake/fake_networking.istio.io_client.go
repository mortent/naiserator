// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha3 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/networking.istio.io/v1alpha3"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeNetworkingV1alpha3 struct {
	*testing.Fake
}

func (c *FakeNetworkingV1alpha3) ServiceEntries(namespace string) v1alpha3.ServiceEntryInterface {
	return &FakeServiceEntries{c, namespace}
}

func (c *FakeNetworkingV1alpha3) VirtualServices(namespace string) v1alpha3.VirtualServiceInterface {
	return &FakeVirtualServices{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNetworkingV1alpha3) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
