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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "agones.dev/agones/pkg/apis/agones/v1"
	scheme "agones.dev/agones/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GameServersGetter has a method to return a GameServerInterface.
// A group's client should implement this interface.
type GameServersGetter interface {
	GameServers(namespace string) GameServerInterface
}

// GameServerInterface has methods to work with GameServer resources.
type GameServerInterface interface {
	Create(ctx context.Context, gameServer *v1.GameServer, opts metav1.CreateOptions) (*v1.GameServer, error)
	Update(ctx context.Context, gameServer *v1.GameServer, opts metav1.UpdateOptions) (*v1.GameServer, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.GameServer, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.GameServerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GameServer, err error)
	GameServerExpansion
}

// gameServers implements GameServerInterface
type gameServers struct {
	client rest.Interface
	ns     string
}

// newGameServers returns a GameServers
func newGameServers(c *AgonesV1Client, namespace string) *gameServers {
	return &gameServers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gameServer, and returns the corresponding gameServer object, and an error if there is any.
func (c *gameServers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.GameServer, err error) {
	result = &v1.GameServer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gameservers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GameServers that match those selectors.
func (c *gameServers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.GameServerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.GameServerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gameservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gameServers.
func (c *gameServers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gameservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a gameServer and creates it.  Returns the server's representation of the gameServer, and an error, if there is any.
func (c *gameServers) Create(ctx context.Context, gameServer *v1.GameServer, opts metav1.CreateOptions) (result *v1.GameServer, err error) {
	result = &v1.GameServer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gameservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gameServer).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a gameServer and updates it. Returns the server's representation of the gameServer, and an error, if there is any.
func (c *gameServers) Update(ctx context.Context, gameServer *v1.GameServer, opts metav1.UpdateOptions) (result *v1.GameServer, err error) {
	result = &v1.GameServer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gameservers").
		Name(gameServer.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gameServer).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the gameServer and deletes it. Returns an error if one occurs.
func (c *gameServers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gameservers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gameServers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gameservers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched gameServer.
func (c *gameServers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GameServer, err error) {
	result = &v1.GameServer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gameservers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
