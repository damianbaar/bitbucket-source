// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/nachocano/bitbucket-source/pkg/apis/sources/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BitBucketSourceLister helps list BitBucketSources.
type BitBucketSourceLister interface {
	// List lists all BitBucketSources in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.BitBucketSource, err error)
	// BitBucketSources returns an object that can list and get BitBucketSources.
	BitBucketSources(namespace string) BitBucketSourceNamespaceLister
	BitBucketSourceListerExpansion
}

// bitBucketSourceLister implements the BitBucketSourceLister interface.
type bitBucketSourceLister struct {
	indexer cache.Indexer
}

// NewBitBucketSourceLister returns a new BitBucketSourceLister.
func NewBitBucketSourceLister(indexer cache.Indexer) BitBucketSourceLister {
	return &bitBucketSourceLister{indexer: indexer}
}

// List lists all BitBucketSources in the indexer.
func (s *bitBucketSourceLister) List(selector labels.Selector) (ret []*v1alpha1.BitBucketSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BitBucketSource))
	})
	return ret, err
}

// BitBucketSources returns an object that can list and get BitBucketSources.
func (s *bitBucketSourceLister) BitBucketSources(namespace string) BitBucketSourceNamespaceLister {
	return bitBucketSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BitBucketSourceNamespaceLister helps list and get BitBucketSources.
type BitBucketSourceNamespaceLister interface {
	// List lists all BitBucketSources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.BitBucketSource, err error)
	// Get retrieves the BitBucketSource from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.BitBucketSource, error)
	BitBucketSourceNamespaceListerExpansion
}

// bitBucketSourceNamespaceLister implements the BitBucketSourceNamespaceLister
// interface.
type bitBucketSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BitBucketSources in the indexer for a given namespace.
func (s bitBucketSourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BitBucketSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BitBucketSource))
	})
	return ret, err
}

// Get retrieves the BitBucketSource from the indexer for a given namespace and name.
func (s bitBucketSourceNamespaceLister) Get(name string) (*v1alpha1.BitBucketSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bitbucketsource"), name)
	}
	return obj.(*v1alpha1.BitBucketSource), nil
}