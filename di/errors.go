package di

import (
	"fmt"
	"reflect"
	"strings"
)

// ErrNotRegistered is returned when attempting to resolve an unregistered type.
//
// This error occurs when you try to resolve a type that has not been registered
// with [Register], [RegisterInstance], or [RegisterType].
//
// Example:
//
//	_, err := di.Resolve[Logger](container)
//	if err != nil {
//	    var notRegistered di.ErrNotRegistered
//	    if errors.As(err, &notRegistered) {
//	        fmt.Printf("Type %s is not registered\n", notRegistered.Type)
//	    }
//	}
type ErrNotRegistered struct {
	// Type is the reflect.Type that was not found in the container.
	Type reflect.Type
}

func (e ErrNotRegistered) Error() string {
	return fmt.Sprintf("di: type %s is not registered", e.Type)
}

// ErrCircularDependency is returned when a circular dependency is detected.
//
// This error occurs when resolving a type would create an infinite loop,
// such as A depends on B, and B depends on A.
//
// The Chain field contains the dependency path that forms the cycle,
// with the repeated type appearing at both the start and end.
//
// Example:
//
//	_, err := di.Resolve[ServiceA](container)
//	if err != nil {
//	    var circular di.ErrCircularDependency
//	    if errors.As(err, &circular) {
//	        fmt.Printf("Circular dependency: %v\n", circular.Chain)
//	        // Output: Circular dependency: [ServiceA ServiceB ServiceA]
//	    }
//	}
type ErrCircularDependency struct {
	// Chain contains the dependency path forming the cycle.
	// The last element is the type that was already being resolved.
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
//
// This wraps the underlying error that caused the resolution to fail.
// Common causes include:
//   - A factory function returned an error
//   - A dependency of the requested type failed to resolve
//
// Use [errors.Unwrap] or the Unwrap method to get the underlying error.
//
// Example:
//
//	_, err := di.Resolve[Database](container)
//	if err != nil {
//	    var resFailed di.ErrResolutionFailed
//	    if errors.As(err, &resFailed) {
//	        fmt.Printf("Failed to resolve %s: %v\n", resFailed.Type, resFailed.Cause)
//	    }
//	}
type ErrResolutionFailed struct {
	// Type is the type that failed to resolve.
	Type reflect.Type
	// Cause is the underlying error that caused the failure.
	Cause error
}

func (e ErrResolutionFailed) Error() string {
	return fmt.Sprintf("di: failed to resolve %s: %v", e.Type, e.Cause)
}

// Unwrap returns the underlying error that caused the resolution failure.
// This allows use with [errors.Is] and [errors.As].
func (e ErrResolutionFailed) Unwrap() error {
	return e.Cause
}

// ErrInvalidFactory is returned when a factory function has an invalid signature.
//
// Valid factory signatures are:
//   - func(...dependencies) T
//   - func(...dependencies) (T, error)
//
// This error occurs when:
//   - The factory is not a function
//   - The factory returns no values
//   - The factory returns more than 2 values
//   - The first return type is not assignable to the registered type
//   - The second return type (if present) is not error
//
// Example:
//
//	err := di.Register[Logger](container, "not a function")
//	if err != nil {
//	    var invalid di.ErrInvalidFactory
//	    if errors.As(err, &invalid) {
//	        fmt.Printf("Invalid factory for %s: %s\n", invalid.Type, invalid.Message)
//	    }
//	}
type ErrInvalidFactory struct {
	// Type is the target type the factory was supposed to create.
	Type reflect.Type
	// Message describes why the factory is invalid.
	Message string
}

func (e ErrInvalidFactory) Error() string {
	return fmt.Sprintf("di: invalid factory for %s: %s", e.Type, e.Message)
}

// ErrScopeNotFound is returned when trying to use a scope that doesn't exist.
//
// This error occurs when attempting to resolve a scoped dependency with a scope
// that hasn't been created via [Container.CreateScope].
type ErrScopeNotFound struct {
	// Name is the name of the scope that was not found.
	Name string
}

func (e ErrScopeNotFound) Error() string {
	return fmt.Sprintf("di: scope %q not found", e.Name)
}
