// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	searchv1alpha1 "github.com/karmada-io/karmada/pkg/apis/search/v1alpha1"
	versioned "github.com/karmada-io/karmada/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/karmada-io/karmada/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/karmada-io/karmada/pkg/generated/listers/search/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ResourceRegistryInformer provides access to a shared informer and lister for
// ResourceRegistries.
type ResourceRegistryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ResourceRegistryLister
}

type resourceRegistryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewResourceRegistryInformer constructs a new informer for ResourceRegistry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceRegistryInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceRegistryInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredResourceRegistryInformer constructs a new informer for ResourceRegistry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceRegistryInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SearchV1alpha1().ResourceRegistries().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SearchV1alpha1().ResourceRegistries().Watch(context.TODO(), options)
			},
		},
		&searchv1alpha1.ResourceRegistry{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceRegistryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceRegistryInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceRegistryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&searchv1alpha1.ResourceRegistry{}, f.defaultInformer)
}

func (f *resourceRegistryInformer) Lister() v1alpha1.ResourceRegistryLister {
	return v1alpha1.NewResourceRegistryLister(f.Informer().GetIndexer())
}
