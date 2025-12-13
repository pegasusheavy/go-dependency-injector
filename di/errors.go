package di

import (
	"fmt"
	"reflect"
	"strings"
)

// ErrNotRegistered is returned when attempting to resolve an unregistered type.
type ErrNotRegistered struct {
	Type reflect.Type
}

func (e ErrNotRegistered) Error() string {
	return fmt.Sprintf("di: type %s is not registered", e.Type)
}

// ErrCircularDependency is returned when a circular dependency is detected.
type ErrCircularDependency struct {
	Chain []reflect.Type
}

func (e ErrCircularDependency) Error() string {
	names := make([]string, len(e.Chain))
	for i, t := range e.Chain {
		names[i] = t.String()
	}
	return fmt.Sprintf("di: circular dependency detected: %s", strings.Join(names, " -> "))
}

// ErrResolutionFailed is returned when dependency resolution fails.
type ErrResolutionFailed struct {
	Type  reflect.Type
	Cause error
}

func (e ErrResolutionFailed) Error() string {
	return fmt.Sprintf("di: failed to resolve %s: %v", e.Type, e.Cause)
}

func (e ErrResolutionFailed) Unwrap() error {
	return e.Cause
}

// ErrInvalidFactory is returned when a factory function has an invalid signature.
type ErrInvalidFactory struct {
	Type    reflect.Type
	Message string
}

func (e ErrInvalidFactory) Error() string {
	return fmt.Sprintf("di: invalid factory for %s: %s", e.Type, e.Message)
}

// ErrScopeNotFound is returned when trying to use a scope that doesn't exist.
type ErrScopeNotFound struct {
	Name string
}

func (e ErrScopeNotFound) Error() string {
	return fmt.Sprintf("di: scope %q not found", e.Name)
}

