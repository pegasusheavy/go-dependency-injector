package di

import (
	"reflect"
	"sync"
)

// Container is the dependency injection container.
type Container struct {
	mu            sync.RWMutex
	registrations map[registrationKey]*registration
	singletons    map[registrationKey]any
	scopes        map[string]*Scope
	resolving     map[reflect.Type]bool // For circular dependency detection
}

// New creates a new dependency injection container.
func New() *Container {
	return &Container{
		registrations: make(map[registrationKey]*registration),
		singletons:    make(map[registrationKey]any),
		scopes:        make(map[string]*Scope),
		resolving:     make(map[reflect.Type]bool),
	}
}

// Register registers a type with the container using a factory function.
// The factory can have dependencies as parameters which will be auto-resolved.
func Register[T any](c *Container, factory any, opts ...RegistrationOption) error {
	var zero T
	targetType := reflect.TypeOf(&zero).Elem()

	reg := &registration{
		targetType: targetType,
		factory:    factory,
		lifetime:   Transient,
	}

	for _, opt := range opts {
		opt(reg)
	}

	// Validate factory signature
	if err := validateFactory(targetType, factory); err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	key := registrationKey{typ: targetType, name: reg.name}
	c.registrations[key] = reg

	return nil
}

// RegisterInstance registers an existing instance as a singleton.
func RegisterInstance[T any](c *Container, instance T, opts ...RegistrationOption) {
	var zero T
	targetType := reflect.TypeOf(&zero).Elem()

	reg := &registration{
		targetType: targetType,
		instance:   instance,
		lifetime:   Singleton,
	}

	for _, opt := range opts {
		opt(reg)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	key := registrationKey{typ: targetType, name: reg.name}
	c.registrations[key] = reg
	c.singletons[key] = instance
}

// RegisterType registers a concrete type that will be instantiated automatically.
// The type must have a constructor function named New<TypeName> or accept struct injection.
func RegisterType[TInterface any, TImpl any](c *Container, opts ...RegistrationOption) error {
	var zeroIface TInterface
	var zeroImpl TImpl
	ifaceType := reflect.TypeOf(&zeroIface).Elem()
	implType := reflect.TypeOf(&zeroImpl).Elem()

	// Create a factory that instantiates the implementation
	factory := func() TInterface {
		val := reflect.New(implType)
		return val.Interface().(TInterface)
	}

	reg := &registration{
		targetType: ifaceType,
		implType:   implType,
		factory:    factory,
		lifetime:   Transient,
	}

	for _, opt := range opts {
		opt(reg)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	key := registrationKey{typ: ifaceType, name: reg.name}
	c.registrations[key] = reg

	return nil
}

// Resolve resolves a dependency from the container.
func Resolve[T any](c *Container) (T, error) {
	return ResolveNamed[T](c, "")
}

// ResolveNamed resolves a named dependency from the container.
func ResolveNamed[T any](c *Container, name string) (T, error) {
	var zero T
	targetType := reflect.TypeOf(&zero).Elem()

	result, err := c.resolve(targetType, name, nil, make([]reflect.Type, 0))
	if err != nil {
		return zero, err
	}

	return result.(T), nil
}

// MustResolve resolves a dependency or panics if it fails.
func MustResolve[T any](c *Container) T {
	result, err := Resolve[T](c)
	if err != nil {
		panic(err)
	}
	return result
}

// MustResolveNamed resolves a named dependency or panics if it fails.
func MustResolveNamed[T any](c *Container, name string) T {
	result, err := ResolveNamed[T](c, name)
	if err != nil {
		panic(err)
	}
	return result
}

// CreateScope creates a new resolution scope.
func (c *Container) CreateScope(name string) *Scope {
	c.mu.Lock()
	defer c.mu.Unlock()

	scope := newScope(name, c)
	c.scopes[name] = scope
	return scope
}

// ResolveInScope resolves a dependency within a specific scope.
func ResolveInScope[T any](c *Container, scope *Scope) (T, error) {
	var zero T
	targetType := reflect.TypeOf(&zero).Elem()

	result, err := c.resolve(targetType, "", scope, make([]reflect.Type, 0))
	if err != nil {
		return zero, err
	}

	return result.(T), nil
}

// resolve is the internal resolution method.
func (c *Container) resolve(targetType reflect.Type, name string, scope *Scope, chain []reflect.Type) (any, error) {
	c.mu.RLock()
	key := registrationKey{typ: targetType, name: name}
	reg, exists := c.registrations[key]
	c.mu.RUnlock()

	if !exists {
		return nil, ErrNotRegistered{Type: targetType}
	}

	// Check for circular dependencies
	for _, t := range chain {
		if t == targetType {
			return nil, ErrCircularDependency{Chain: append(chain, targetType)}
		}
	}
	chain = append(chain, targetType)

	// Handle pre-registered instances
	if reg.instance != nil {
		return reg.instance, nil
	}

	// Check singleton cache
	if reg.lifetime == Singleton {
		c.mu.RLock()
		if instance, ok := c.singletons[key]; ok {
			c.mu.RUnlock()
			return instance, nil
		}
		c.mu.RUnlock()
	}

	// Check scope cache for scoped dependencies
	if reg.lifetime == Scoped && scope != nil {
		if instance, ok := scope.get(key); ok {
			return instance, nil
		}
	}

	// Create new instance using factory
	instance, err := c.invokeFactory(reg.factory, scope, chain)
	if err != nil {
		return nil, ErrResolutionFailed{Type: targetType, Cause: err}
	}

	// Cache based on lifetime
	switch reg.lifetime {
	case Singleton:
		c.mu.Lock()
		c.singletons[key] = instance
		c.mu.Unlock()
	case Scoped:
		if scope != nil {
			scope.set(key, instance)
		}
	}

	return instance, nil
}

// invokeFactory calls a factory function, resolving its dependencies.
func (c *Container) invokeFactory(factory any, scope *Scope, chain []reflect.Type) (any, error) {
	factoryValue := reflect.ValueOf(factory)
	factoryType := factoryValue.Type()

	// Resolve all parameters
	args := make([]reflect.Value, factoryType.NumIn())
	for i := 0; i < factoryType.NumIn(); i++ {
		paramType := factoryType.In(i)
		resolved, err := c.resolve(paramType, "", scope, chain)
		if err != nil {
			return nil, err
		}
		args[i] = reflect.ValueOf(resolved)
	}

	// Call factory
	results := factoryValue.Call(args)

	// Handle (T) or (T, error) return signatures
	if len(results) == 0 {
		return nil, ErrInvalidFactory{Type: factoryType, Message: "factory must return at least one value"}
	}

	// Check for error return
	if len(results) == 2 {
		if !results[1].IsNil() {
			return nil, results[1].Interface().(error)
		}
	}

	return results[0].Interface(), nil
}

// validateFactory ensures the factory has a valid signature.
func validateFactory(targetType reflect.Type, factory any) error {
	factoryValue := reflect.ValueOf(factory)
	if factoryValue.Kind() != reflect.Func {
		return ErrInvalidFactory{Type: targetType, Message: "factory must be a function"}
	}

	factoryType := factoryValue.Type()

	// Must return at least one value
	if factoryType.NumOut() == 0 {
		return ErrInvalidFactory{Type: targetType, Message: "factory must return a value"}
	}

	// First return type must be assignable to target type
	returnType := factoryType.Out(0)
	if !returnType.AssignableTo(targetType) && !(targetType.Kind() == reflect.Interface && returnType.Implements(targetType)) {
		return ErrInvalidFactory{
			Type:    targetType,
			Message: "factory return type " + returnType.String() + " is not assignable to " + targetType.String(),
		}
	}

	// If two return values, second must be error
	if factoryType.NumOut() == 2 {
		errorType := reflect.TypeOf((*error)(nil)).Elem()
		if !factoryType.Out(1).Implements(errorType) {
			return ErrInvalidFactory{Type: targetType, Message: "second return value must be error"}
		}
	}

	// Cannot have more than 2 return values
	if factoryType.NumOut() > 2 {
		return ErrInvalidFactory{Type: targetType, Message: "factory cannot return more than 2 values"}
	}

	return nil
}

// Has checks if a type is registered.
func Has[T any](c *Container) bool {
	return HasNamed[T](c, "")
}

// HasNamed checks if a named type is registered.
func HasNamed[T any](c *Container, name string) bool {
	var zero T
	targetType := reflect.TypeOf(&zero).Elem()

	c.mu.RLock()
	defer c.mu.RUnlock()

	key := registrationKey{typ: targetType, name: name}
	_, exists := c.registrations[key]
	return exists
}

// Clear removes all registrations and cached instances.
func (c *Container) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.registrations = make(map[registrationKey]*registration)
	c.singletons = make(map[registrationKey]any)
	c.scopes = make(map[string]*Scope)
}
