// Copyright 2022 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "agones.dev/agones/pkg/apis/agones/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FleetLister helps list Fleets.
// All objects returned here must be treated as read-only.
type FleetLister interface {
	// List lists all Fleets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Fleet, err error)
	// Fleets returns an object that can list and get Fleets.
	Fleets(namespace string) FleetNamespaceLister
	FleetListerExpansion
}

// fleetLister implements the FleetLister interface.
type fleetLister struct {
	indexer cache.Indexer
}

// NewFleetLister returns a new FleetLister.
func NewFleetLister(indexer cache.Indexer) FleetLister {
	return &fleetLister{indexer: indexer}
}

// List lists all Fleets in the indexer.
func (s *fleetLister) List(selector labels.Selector) (ret []*v1.Fleet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Fleet))
	})
	return ret, err
}

// Fleets returns an object that can list and get Fleets.
func (s *fleetLister) Fleets(namespace string) FleetNamespaceLister {
	return fleetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FleetNamespaceLister helps list and get Fleets.
// All objects returned here must be treated as read-only.
type FleetNamespaceLister interface {
	// List lists all Fleets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Fleet, err error)
	// Get retrieves the Fleet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Fleet, error)
	FleetNamespaceListerExpansion
}

// fleetNamespaceLister implements the FleetNamespaceLister
// interface.
type fleetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Fleets in the indexer for a given namespace.
func (s fleetNamespaceLister) List(selector labels.Selector) (ret []*v1.Fleet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Fleet))
	})
	return ret, err
}

// Get retrieves the Fleet from the indexer for a given namespace and name.
func (s fleetNamespaceLister) Get(name string) (*v1.Fleet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("fleet"), name)
	}
	return obj.(*v1.Fleet), nil
}
