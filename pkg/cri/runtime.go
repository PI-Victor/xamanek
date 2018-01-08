// Package runtime abstract away the implementation of the current container
// runtime to make it possible to swap them.
package cri

type ContainerRuntimeClient interface{}

type ContainerRuntimeInterface interface {
	CreateContainer() error
	DestroyContainer() error
}
