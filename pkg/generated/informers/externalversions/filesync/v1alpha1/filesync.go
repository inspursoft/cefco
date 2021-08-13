/*
Copyright 2021 The idx Authors.

http://git.inspur.com/middleware/idx-component

### It needs to be supplemented ###
### It needs to be supplemented ###
### It needs to be supplemented ###
### It needs to be supplemented ###
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	filesyncv1alpha1 "git.inspur.com/middleware/idx-component/docker/filesync/idx-ceco/pkg/apis/filesync/v1alpha1"
	versioned "git.inspur.com/middleware/idx-component/docker/filesync/idx-ceco/pkg/generated/clientset/versioned"
	internalinterfaces "git.inspur.com/middleware/idx-component/docker/filesync/idx-ceco/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "git.inspur.com/middleware/idx-component/docker/filesync/idx-ceco/pkg/generated/listers/filesync/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FileSyncInformer provides access to a shared informer and lister for
// FileSyncs.
type FileSyncInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FileSyncLister
}

type fileSyncInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFileSyncInformer constructs a new informer for FileSync type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFileSyncInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFileSyncInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFileSyncInformer constructs a new informer for FileSync type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFileSyncInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IdxV1alpha1().FileSyncs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IdxV1alpha1().FileSyncs(namespace).Watch(context.TODO(), options)
			},
		},
		&filesyncv1alpha1.FileSync{},
		resyncPeriod,
		indexers,
	)
}

func (f *fileSyncInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFileSyncInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fileSyncInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&filesyncv1alpha1.FileSync{}, f.defaultInformer)
}

func (f *fileSyncInformer) Lister() v1alpha1.FileSyncLister {
	return v1alpha1.NewFileSyncLister(f.Informer().GetIndexer())
}