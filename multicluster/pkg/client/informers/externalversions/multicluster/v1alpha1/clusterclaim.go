/*
Copyright 2021 Antrea Authors.

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
	"context"
	time "time"

	multiclusterv1alpha1 "antrea.io/antrea/multicluster/apis/multicluster/v1alpha1"
	versioned "antrea.io/antrea/multicluster/pkg/client/clientset/versioned"
	internalinterfaces "antrea.io/antrea/multicluster/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "antrea.io/antrea/multicluster/pkg/client/listers/multicluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterClaimInformer provides access to a shared informer and lister for
// ClusterClaims.
type ClusterClaimInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterClaimLister
}

type clusterClaimInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterClaimInformer constructs a new informer for ClusterClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterClaimInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterClaimInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterClaimInformer constructs a new informer for ClusterClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterClaimInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MulticlusterV1alpha1().ClusterClaims(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MulticlusterV1alpha1().ClusterClaims(namespace).Watch(context.TODO(), options)
			},
		},
		&multiclusterv1alpha1.ClusterClaim{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterClaimInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterClaimInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterClaimInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&multiclusterv1alpha1.ClusterClaim{}, f.defaultInformer)
}

func (f *clusterClaimInformer) Lister() v1alpha1.ClusterClaimLister {
	return v1alpha1.NewClusterClaimLister(f.Informer().GetIndexer())
}