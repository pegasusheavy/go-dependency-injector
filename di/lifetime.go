package di

import "sync"

// Lifetime defines how long a resolved instance lives.
//
// The lifetime determines caching behavior:
//   - [Transient]: No caching, new instance every resolution
//   - [Singleton]: Cached for container lifetime
//   - [Scoped]: Cached per scope
//
// Use the [WithLifetime] option or convenience functions [AsSingleton],
// [AsTransient], [AsScoped] when registering dependencies.
type Lifetime int

const (
	// Transient creates a new instance every time it's resolved.
	// This is the default lifetime.
	//
	// Use for stateful objects or when each consumer needs its own instance.
	//
	// Example:
	//
	//	di.Register[Handler](c, factory) // default is Transient
	//	di.Register[Handler](c, factory, di.AsTransient())
	Transient Lifetime = iota

	// Singleton creates one instance for the entire container lifetime.
	// The instance is created on first resolution and cached.
	//
	// Use for stateless services, configuration, and shared resources.
	//
	// Example:
	//
	//	di.Register[Logger](c, factory, di.AsSingleton())
	Singleton

	// Scoped creates one instance per scope.
	// Within a scope, the same instance is returned.
	// Different scopes get different instances.
	//
	// Use for request-scoped dependencies in web applications.
	//
	// Example:
	//
	//	di.Register[*RequestContext](c, factory, di.AsScoped())
	//
	//	scope := container.CreateScope("request-1")
	//	ctx, _ := di.ResolveInScope[*RequestContext](c, scope)
	Scoped
)

// String returns the string representation of the lifetime.
func (l Lifetime) String() string {
	switch l {
	case Transient:
		return "Transient"
	case Singleton:
		return "Singleton"
	case Scoped:
		return "Scoped"
	default:
		return "Unknown"
	}
}

// Scope represents a resolution scope for scoped dependencies.
//
// Scopes are created via [Container.CreateScope] and used with [ResolveInScope]
// to resolve scoped dependencies. Within a scope, scoped dependencies return
// the same instance.
//
// Scopes are useful for request-scoped dependencies in web applications,
// where you want the same instance throughout a single HTTP request but
// different instances for different requests.
//
// Example:
//
//	// Per HTTP request
//	func handler(w http.ResponseWriter, r *http.Request) {
//	    scope := container.CreateScope("request-" + requestID)
//
//	    // Same instance within this request
//	    ctx1, _ := di.ResolveInScope[*RequestContext](c, scope)
//	    ctx2, _ := di.ResolveInScope[*RequestContext](c, scope)
//	    // ctx1 == ctx2
//	}
type Scope struct {
	mu        sync.RWMutex
	name      string
	instances map[any]any
	parent    *Container
}

// newScope creates a new scope attached to the given container.
func newScope(name string, parent *Container) *Scope {
	return &Scope{
		name:      name,
		instances: make(map[any]any),
		parent:    parent,
	}
}

// Name returns the scope's identifier.
//
// This is the name that was passed to [Container.CreateScope].
func (s *Scope) Name() string {
	return s.name
}

// get retrieves an instance from the scope cache.
func (s *Scope) get(key any) (any, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	instance, ok := s.instances[key]
	return instance, ok
}

// set stores an instance in the scope cache.
func (s *Scope) set(key any, instance any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.instances[key] = instance
}
