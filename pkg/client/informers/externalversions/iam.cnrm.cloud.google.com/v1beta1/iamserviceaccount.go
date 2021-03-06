// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	iamcnrmcloudgooglecomv1beta1 "github.com/nais/naiserator/pkg/apis/iam.cnrm.cloud.google.com/v1beta1"
	versioned "github.com/nais/naiserator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nais/naiserator/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/nais/naiserator/pkg/client/listers/iam.cnrm.cloud.google.com/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// IAMServiceAccountInformer provides access to a shared informer and lister for
// IAMServiceAccounts.
type IAMServiceAccountInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.IAMServiceAccountLister
}

type iAMServiceAccountInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewIAMServiceAccountInformer constructs a new informer for IAMServiceAccount type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIAMServiceAccountInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIAMServiceAccountInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredIAMServiceAccountInformer constructs a new informer for IAMServiceAccount type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIAMServiceAccountInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IamV1beta1().IAMServiceAccounts(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IamV1beta1().IAMServiceAccounts(namespace).Watch(options)
			},
		},
		&iamcnrmcloudgooglecomv1beta1.IAMServiceAccount{},
		resyncPeriod,
		indexers,
	)
}

func (f *iAMServiceAccountInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIAMServiceAccountInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *iAMServiceAccountInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&iamcnrmcloudgooglecomv1beta1.IAMServiceAccount{}, f.defaultInformer)
}

func (f *iAMServiceAccountInformer) Lister() v1beta1.IAMServiceAccountLister {
	return v1beta1.NewIAMServiceAccountLister(f.Informer().GetIndexer())
}
