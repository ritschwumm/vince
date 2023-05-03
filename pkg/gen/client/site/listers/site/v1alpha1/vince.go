/*
Licensed under the GNU AFFERO GENERAL PUBLIC LICENSE Version 3
*/
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gernest/vince/pkg/apis/site/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VinceLister helps list Vinces.
// All objects returned here must be treated as read-only.
type VinceLister interface {
	// List lists all Vinces in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Vince, err error)
	// Vinces returns an object that can list and get Vinces.
	Vinces(namespace string) VinceNamespaceLister
	VinceListerExpansion
}

// vinceLister implements the VinceLister interface.
type vinceLister struct {
	indexer cache.Indexer
}

// NewVinceLister returns a new VinceLister.
func NewVinceLister(indexer cache.Indexer) VinceLister {
	return &vinceLister{indexer: indexer}
}

// List lists all Vinces in the indexer.
func (s *vinceLister) List(selector labels.Selector) (ret []*v1alpha1.Vince, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Vince))
	})
	return ret, err
}

// Vinces returns an object that can list and get Vinces.
func (s *vinceLister) Vinces(namespace string) VinceNamespaceLister {
	return vinceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VinceNamespaceLister helps list and get Vinces.
// All objects returned here must be treated as read-only.
type VinceNamespaceLister interface {
	// List lists all Vinces in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Vince, err error)
	// Get retrieves the Vince from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Vince, error)
	VinceNamespaceListerExpansion
}

// vinceNamespaceLister implements the VinceNamespaceLister
// interface.
type vinceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Vinces in the indexer for a given namespace.
func (s vinceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Vince, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Vince))
	})
	return ret, err
}

// Get retrieves the Vince from the indexer for a given namespace and name.
func (s vinceNamespaceLister) Get(name string) (*v1alpha1.Vince, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vince"), name)
	}
	return obj.(*v1alpha1.Vince), nil
}
