// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/nais/naiserator/pkg/apis/iam.cnrm.cloud.google.com/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IAMPolicyLister helps list IAMPolicies.
type IAMPolicyLister interface {
	// List lists all IAMPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IAMPolicy, err error)
	// IAMPolicies returns an object that can list and get IAMPolicies.
	IAMPolicies(namespace string) IAMPolicyNamespaceLister
	IAMPolicyListerExpansion
}

// iAMPolicyLister implements the IAMPolicyLister interface.
type iAMPolicyLister struct {
	indexer cache.Indexer
}

// NewIAMPolicyLister returns a new IAMPolicyLister.
func NewIAMPolicyLister(indexer cache.Indexer) IAMPolicyLister {
	return &iAMPolicyLister{indexer: indexer}
}

// List lists all IAMPolicies in the indexer.
func (s *iAMPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.IAMPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IAMPolicy))
	})
	return ret, err
}

// IAMPolicies returns an object that can list and get IAMPolicies.
func (s *iAMPolicyLister) IAMPolicies(namespace string) IAMPolicyNamespaceLister {
	return iAMPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IAMPolicyNamespaceLister helps list and get IAMPolicies.
type IAMPolicyNamespaceLister interface {
	// List lists all IAMPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IAMPolicy, err error)
	// Get retrieves the IAMPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IAMPolicy, error)
	IAMPolicyNamespaceListerExpansion
}

// iAMPolicyNamespaceLister implements the IAMPolicyNamespaceLister
// interface.
type iAMPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IAMPolicies in the indexer for a given namespace.
func (s iAMPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IAMPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IAMPolicy))
	})
	return ret, err
}

// Get retrieves the IAMPolicy from the indexer for a given namespace and name.
func (s iAMPolicyNamespaceLister) Get(name string) (*v1alpha1.IAMPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iampolicy"), name)
	}
	return obj.(*v1alpha1.IAMPolicy), nil
}