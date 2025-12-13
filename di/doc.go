// Package di provides a lightweight, type-safe dependency injection container for Go.
//
// # Features
//
//   - Generic type-safe registration and resolution
//   - Multiple lifetimes: Transient, Singleton, and Scoped
//   - Automatic dependency resolution via constructor injection
//   - Circular dependency detection
//   - Named registrations for multiple implementations
//   - Thread-safe operations
//
// # Basic Usage
//
//	container := di.New()
//
//	// Register with factory
//	di.Register[Logger](container, func() Logger {
//	    return &ConsoleLogger{}
//	}, di.AsSingleton())
//
//	// Register with dependencies
//	di.Register[UserService](container, func(log Logger) UserService {
//	    return &DefaultUserService{logger: log}
//	})
//
//	// Resolve
//	service, err := di.Resolve[UserService](container)
//
// # Lifetimes
//
// Transient: A new instance is created every time the dependency is resolved.
// This is the default lifetime.
//
// Singleton: A single instance is created and reused for all resolutions.
//
// Scoped: A single instance per scope. Useful for request-scoped dependencies
// in web applications.
package di

