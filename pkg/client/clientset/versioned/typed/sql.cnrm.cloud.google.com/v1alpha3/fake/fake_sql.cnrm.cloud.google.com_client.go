// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha3 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/sql.cnrm.cloud.google.com/v1alpha3"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSqlV1alpha3 struct {
	*testing.Fake
}

func (c *FakeSqlV1alpha3) SQLDatabases(namespace string) v1alpha3.SQLDatabaseInterface {
	return &FakeSQLDatabases{c, namespace}
}

func (c *FakeSqlV1alpha3) SQLInstances(namespace string) v1alpha3.SQLInstanceInterface {
	return &FakeSQLInstances{c, namespace}
}

func (c *FakeSqlV1alpha3) SQLUsers(namespace string) v1alpha3.SQLUserInterface {
	return &FakeSQLUsers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSqlV1alpha3) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
