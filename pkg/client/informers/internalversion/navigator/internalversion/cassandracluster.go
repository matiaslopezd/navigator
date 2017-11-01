/*
Copyright 2017 Jetstack Ltd.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	navigator "github.com/jetstack/navigator/pkg/apis/navigator"
	clientset_internalversion "github.com/jetstack/navigator/pkg/client/clientset/internalversion"
	internalinterfaces "github.com/jetstack/navigator/pkg/client/informers/internalversion/internalinterfaces"
	internalversion "github.com/jetstack/navigator/pkg/client/listers/navigator/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// CassandraClusterInformer provides access to a shared informer and lister for
// CassandraClusters.
type CassandraClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.CassandraClusterLister
}

type cassandraClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	filter  internalinterfaces.FilterFunc
}

// NewCassandraClusterInformer constructs a new informer for CassandraCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCassandraClusterInformer(client clientset_internalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	filter := internalinterfaces.NamespaceFilter(namespace)
	return NewFilteredCassandraClusterInformer(client, filter, resyncPeriod, indexers)
}

// NewFilteredCassandraClusterInformer constructs a new informer for CassandraCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCassandraClusterInformer(client clientset_internalversion.Interface, filter internalinterfaces.FilterFunc, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				namespace := filter(&options)
				return client.Navigator().CassandraClusters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				namespace := filter(&options)
				return client.Navigator().CassandraClusters(namespace).Watch(options)
			},
		},
		&navigator.CassandraCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *cassandraClusterInformer) defaultInformer(client clientset_internalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCassandraClusterInformer(client, f.filter, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *cassandraClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&navigator.CassandraCluster{}, f.defaultInformer)
}

func (f *cassandraClusterInformer) Lister() internalversion.CassandraClusterLister {
	return internalversion.NewCassandraClusterLister(f.Informer().GetIndexer())
}
