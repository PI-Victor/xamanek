// Package runtime abstract away the implementation of the current container
// runtime to make it possible to swap them.
package runtime

type ContainerRuntimeClient interface{}

type ContainerRuntimeInterface interface {
	CreateContainer() (*ContainerRuntimeClient, error)
	DestroyContainer() error
	ScaleContainers() error
}
