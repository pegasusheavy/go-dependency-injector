package di

import "reflect"

// registration holds metadata about a registered dependency.
type registration struct {
	// The type being registered (interface or concrete type).
	targetType reflect.Type

	// The concrete implementation type.
	implType reflect.Type

	// Factory function to create instances.
	factory any

	// Lifetime of the dependency.
	lifetime Lifetime

	// Pre-created instance (for RegisterInstance).
	instance any

	// Name for named registrations.
	name string
}

// RegistrationOption configures a registration.
type RegistrationOption func(*registration)

// WithLifetime sets the lifetime for the registration.
func WithLifetime(lifetime Lifetime) RegistrationOption {
	return func(r *registration) {
		r.lifetime = lifetime
	}
}

// AsSingleton registers the dependency as a singleton.
func AsSingleton() RegistrationOption {
	return func(r *registration) {
		r.lifetime = Singleton
	}
}

// AsTransient registers the dependency as transient.
func AsTransient() RegistrationOption {
	return func(r *registration) {
		r.lifetime = Transient
	}
}

// AsScoped registers the dependency as scoped.
func AsScoped() RegistrationOption {
	return func(r *registration) {
		r.lifetime = Scoped
	}
}

// WithName sets a name for named registrations.
func WithName(name string) RegistrationOption {
	return func(r *registration) {
		r.name = name
	}
}

// registrationKey uniquely identifies a registration.
type registrationKey struct {
	typ  reflect.Type
	name string
}

