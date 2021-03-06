/*
Copyright 2019 the original author or authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	bindingsv1alpha1 "github.com/projectriff/bindings/pkg/apis/bindings/v1alpha1"
	versioned "github.com/projectriff/bindings/pkg/client/clientset/versioned"
	internalinterfaces "github.com/projectriff/bindings/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/projectriff/bindings/pkg/client/listers/bindings/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ImageBindingInformer provides access to a shared informer and lister for
// ImageBindings.
type ImageBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ImageBindingLister
}

type imageBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewImageBindingInformer constructs a new informer for ImageBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewImageBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredImageBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredImageBindingInformer constructs a new informer for ImageBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredImageBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BindingsV1alpha1().ImageBindings(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BindingsV1alpha1().ImageBindings(namespace).Watch(options)
			},
		},
		&bindingsv1alpha1.ImageBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *imageBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredImageBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *imageBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&bindingsv1alpha1.ImageBinding{}, f.defaultInformer)
}

func (f *imageBindingInformer) Lister() v1alpha1.ImageBindingLister {
	return v1alpha1.NewImageBindingLister(f.Informer().GetIndexer())
}
