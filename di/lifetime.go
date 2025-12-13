package di

import "sync"

// Lifetime defines how long a resolved instance lives.
type Lifetime int

const (
	// Transient creates a new instance every time it's resolved.
	Transient Lifetime = iota

	// Singleton creates one instance for the entire container lifetime.
	Singleton

	// Scoped creates one instance per scope (useful for request-scoped dependencies).
	Scoped
)

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
type Scope struct {
	mu        sync.RWMutex
	name      string
	instances map[any]any
	parent    *Container
}

// newScope creates a new scope.
func newScope(name string, parent *Container) *Scope {
	return &Scope{
		name:      name,
		instances: make(map[any]any),
		parent:    parent,
	}
}

// Name returns the scope's name.
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
