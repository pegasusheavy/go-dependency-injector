package di

import "reflect"

// registration holds metadata about a registered dependency.
// This is an internal type used by the container.
type registration struct {
	// targetType is the type being registered (interface or concrete type).
	targetType reflect.Type

	// implType is the concrete implementation type (used by RegisterType).
	implType reflect.Type

	// factory is the function to create instances.
	factory any

	// lifetime determines how long resolved instances live.
	lifetime Lifetime

	// instance is a pre-created instance (used by RegisterInstance).
	instance any

	// name is the identifier for named registrations.
	name string
}

// RegistrationOption configures a dependency registration.
//
// Options are passed to [Register], [RegisterInstance], and [RegisterType]
// to customize the registration behavior.
//
// Available options:
//   - [AsSingleton]: Single instance shared across all resolutions
//   - [AsTransient]: New instance on each resolution (default)
//   - [AsScoped]: Single instance per scope
//   - [WithLifetime]: Set lifetime explicitly
//   - [WithName]: Register with a name for named resolution
type RegistrationOption func(*registration)

// WithLifetime sets the lifetime for the registration.
//
// This is an explicit way to set the lifetime. You can also use the
// convenience functions [AsSingleton], [AsTransient], or [AsScoped].
//
// Example:
//
//	di.Register[Logger](c, factory, di.WithLifetime(di.Singleton))
func WithLifetime(lifetime Lifetime) RegistrationOption {
	return func(r *registration) {
		r.lifetime = lifetime
	}
}

// AsSingleton registers the dependency as a singleton.
//
// Singleton dependencies are created once and the same instance is returned
// for all subsequent resolutions. Use for:
//   - Stateless services
//   - Configuration objects
//   - Connection pools and shared resources
//
// Example:
//
//	di.Register[Logger](c, func() Logger {
//	    return &ConsoleLogger{}
//	}, di.AsSingleton())
func AsSingleton() RegistrationOption {
	return func(r *registration) {
		r.lifetime = Singleton
	}
}

// AsTransient registers the dependency as transient (default).
//
// Transient dependencies create a new instance every time they are resolved.
// This is the default lifetime if no option is specified. Use for:
//   - Stateful objects
//   - Objects that should not be shared
//   - Lightweight objects with no shared state
//
// Example:
//
//	di.Register[RequestHandler](c, func() RequestHandler {
//	    return &MyHandler{startTime: time.Now()}
//	}, di.AsTransient())
func AsTransient() RegistrationOption {
	return func(r *registration) {
		r.lifetime = Transient
	}
}

// AsScoped registers the dependency as scoped.
//
// Scoped dependencies create one instance per scope. Within a scope,
// the same instance is returned. Different scopes get different instances.
// Use for:
//   - Request-scoped dependencies in web applications
//   - Unit of work patterns
//   - Per-operation context
//
// Example:
//
//	di.Register[*RequestContext](c, func() *RequestContext {
//	    return &RequestContext{ID: uuid.New()}
//	}, di.AsScoped())
//
//	// In handler:
//	scope := container.CreateScope("request-123")
//	ctx, _ := di.ResolveInScope[*RequestContext](container, scope)
func AsScoped() RegistrationOption {
	return func(r *registration) {
		r.lifetime = Scoped
	}
}

// WithName sets a name for named registrations.
//
// Named registrations allow multiple implementations of the same interface
// to be registered and resolved by name. Use [ResolveNamed] or [MustResolveNamed]
// to resolve named dependencies.
//
// Example:
//
//	// Register multiple loggers
//	di.Register[Logger](c, consoleFactory, di.WithName("console"))
//	di.Register[Logger](c, fileFactory, di.WithName("file"))
//	di.Register[Logger](c, jsonFactory, di.WithName("json"))
//
//	// Resolve by name
//	consoleLogger, _ := di.ResolveNamed[Logger](c, "console")
func WithName(name string) RegistrationOption {
	return func(r *registration) {
		r.name = name
	}
}

// registrationKey uniquely identifies a registration by type and optional name.
type registrationKey struct {
	typ  reflect.Type
	name string
}
