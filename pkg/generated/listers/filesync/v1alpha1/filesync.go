/*
Copyright 2021 The idx Authors.

http://git.inspur.com/middleware/idx-component

### It needs to be supplemented ###
### It needs to be supplemented ###
### It needs to be supplemented ###
### It needs to be supplemented ###
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/inspursoft/cefco/pkg/apis/filesync/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FileSyncLister helps list FileSyncs.
// All objects returned here must be treated as read-only.
type FileSyncLister interface {
	// List lists all FileSyncs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FileSync, err error)
	// FileSyncs returns an object that can list and get FileSyncs.
	FileSyncs(namespace string) FileSyncNamespaceLister
	FileSyncListerExpansion
}

// fileSyncLister implements the FileSyncLister interface.
type fileSyncLister struct {
	indexer cache.Indexer
}

// NewFileSyncLister returns a new FileSyncLister.
func NewFileSyncLister(indexer cache.Indexer) FileSyncLister {
	return &fileSyncLister{indexer: indexer}
}

// List lists all FileSyncs in the indexer.
func (s *fileSyncLister) List(selector labels.Selector) (ret []*v1alpha1.FileSync, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FileSync))
	})
	return ret, err
}

// FileSyncs returns an object that can list and get FileSyncs.
func (s *fileSyncLister) FileSyncs(namespace string) FileSyncNamespaceLister {
	return fileSyncNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FileSyncNamespaceLister helps list and get FileSyncs.
// All objects returned here must be treated as read-only.
type FileSyncNamespaceLister interface {
	// List lists all FileSyncs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FileSync, err error)
	// Get retrieves the FileSync from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FileSync, error)
	FileSyncNamespaceListerExpansion
}

// fileSyncNamespaceLister implements the FileSyncNamespaceLister
// interface.
type fileSyncNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FileSyncs in the indexer for a given namespace.
func (s fileSyncNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FileSync, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FileSync))
	})
	return ret, err
}

// Get retrieves the FileSync from the indexer for a given namespace and name.
func (s fileSyncNamespaceLister) Get(name string) (*v1alpha1.FileSync, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("filesync"), name)
	}
	return obj.(*v1alpha1.FileSync), nil
}
