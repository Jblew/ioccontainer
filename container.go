// Package container provides an IoC container for Go projects.
// It provides simple, fluent and easy-to-use interface to make dependency injection in GoLang easier.
package ioccontainer

import (
	internal "github.com/Jblew/ioccontainer/pkg/ioccontainer"
)

// NewContainer makes new container
func NewContainer() *internal.Container {
	return &internal.Container{}
}

// A default instance for container
var container *internal.Container = NewContainer()

// Singleton creates a singleton for the default instance.
func Singleton(resolver interface{}) {
	container.Singleton(resolver)
}

// Transient creates a transient binding for the default instance.
func Transient(resolver interface{}) {
	container.Transient(resolver)
}

// Reset removes all bindings in the default instance.
func Reset() {
	container.Reset()
}

// Make binds receiver to the default instance.
func Make(receiver interface{}) {
	container.Make(receiver)
}
