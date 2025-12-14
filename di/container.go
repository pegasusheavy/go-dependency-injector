package di

import (
	"reflect"
	"sync"
)

// Container is the dependency injection container that manages service registrations
// and their resolution. It is thread-safe and can be used concurrently from multiple
// goroutines.
//
// A Container maintains:
//   - Registrations: Factory functions and metadata for creating instances
//   - Singletons: Cached instances for singleton-scoped dependencies
//   - Scopes: Named scopes for scoped dependency resolution
//
// Use [New] to create a new Container instance.
type Container struct {
	mu            sync.RWMutex
	registrations map[registrationKey]*registration
	singletons    map[registrationKey]any
	scopes        map[string]*Scope
	resolving     map[reflect.Type]bool // For circular dependency detection
}

// New creates a new dependency injection container.
//
// The returned container is empty and ready for registrations. It is thread-safe
// and can be safely shared across goroutines.
//
// Example:
//
//	container := di.New()
//	di.Register[Logger](container, func() Logger { return &ConsoleLogger{} })
func New() *Container {
	return &Container{
		registrations: make(map[registrationKey]*registration),
		singletons:    make(map[registrationKey]any),
		scopes:        make(map[string]*Scope),
		resolving:     make(map[reflect.Type]bool),
	}
}

// Register registers a type with the container using a factory function.
//
// The factory function can take any number of parameters, which will be automatically
// resolved from the container when the type is resolved. The factory must return
// either a single value of type T, or (T, error) if initialization can fail.
//
// By default, registrations are transient (a new instance is created on each resolution).
// Use [AsSingleton], [AsScoped], or [WithLifetime] options to change the lifetime.
//
// Returns an error if the factory signature is invalid (see [ErrInvalidFactory]).
//
// Example:
//
//	// Simple factory with no dependencies
//	di.Register[Logger](c, func() Logger {
//	    return &ConsoleLogger{}
//	})
//
//	// Factory with auto-resolved dependencies
//	di.Register[UserService](c, func(log Logger, repo UserRepository) UserService {
//	    return &DefaultUserService{logger: log, repo: repo}
//	})
//
//	// Factory that can return an error
//	di.Register[*sql.DB](c, func(cfg Config) (*sql.DB, error) {
//	    return sql.Open("postgres", cfg.DatabaseURL())
//	}, di.AsSingleton())
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
//
// Use this when you have a pre-created object that should be returned
// whenever the type is resolved. The instance is always treated as a singleton.
//
// This is useful for:
//   - Configuration objects created at startup
//   - Shared resources like connection pools
//   - Mock objects in tests
//
// Example:
//
//	config := &AppConfig{Port: 8080, Debug: true}
//	di.RegisterInstance[Config](container, config)
//
//	// Later, resolving returns the same instance
//	cfg := di.MustResolve[Config](container)
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

// RegisterType registers an interface to implementation type mapping.
//
// This creates a registration where resolving TInterface returns a new instance
// of TImpl. The implementation type is instantiated using reflection.
//
// This is useful when the implementation has no constructor dependencies and
// can be zero-value initialized, or when you want the container to create
// instances automatically.
//
// Example:
//
//	// Register Logger interface to resolve as ConsoleLogger
//	di.RegisterType[Logger, ConsoleLogger](container, di.AsSingleton())
//
//	// Now resolving Logger returns a *ConsoleLogger
//	logger := di.MustResolve[Logger](container)
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
//
// This returns the resolved instance of type T, or an error if resolution fails.
// All dependencies of T are automatically resolved recursively.
//
// Possible errors:
//   - [ErrNotRegistered]: The type T was not registered
//   - [ErrCircularDependency]: A circular dependency was detected
//   - [ErrResolutionFailed]: The factory or a dependency failed
//
// Example:
//
//	service, err := di.Resolve[UserService](container)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	service.DoSomething()
func Resolve[T any](c *Container) (T, error) {
	return ResolveNamed[T](c, "")
}

// ResolveNamed resolves a named dependency from the container.
//
// Use this when multiple implementations of the same interface are registered
// with different names using [WithName].
//
// Example:
//
//	// Registration
//	di.Register[Logger](c, consoleFactory, di.WithName("console"))
//	di.Register[Logger](c, fileFactory, di.WithName("file"))
//
//	// Resolution by name
//	consoleLogger, err := di.ResolveNamed[Logger](c, "console")
//	fileLogger, err := di.ResolveNamed[Logger](c, "file")
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
//
// Use this when you're certain the type is registered and resolution will succeed,
// such as during application startup after all registrations are complete.
//
// This is equivalent to:
//
//	result, err := di.Resolve[T](c)
//	if err != nil {
//	    panic(err)
//	}
//	return result
//
// Example:
//
//	// Safe to use after startup when you know types are registered
//	service := di.MustResolve[UserService](container)
func MustResolve[T any](c *Container) T {
	result, err := Resolve[T](c)
	if err != nil {
		panic(err)
	}
	return result
}

// MustResolveNamed resolves a named dependency or panics if it fails.
//
// This is the panic-on-error variant of [ResolveNamed]. Use when you're certain
// the named registration exists.
//
// Example:
//
//	// Safe to use when you know the named registration exists
//	logger := di.MustResolveNamed[Logger](container, "console")
func MustResolveNamed[T any](c *Container, name string) T {
	result, err := ResolveNamed[T](c, name)
	if err != nil {
		panic(err)
	}
	return result
}

// CreateScope creates a new resolution scope for scoped dependencies.
//
// Scopes are useful for request-scoped dependencies in web applications.
// Dependencies registered with [AsScoped] will return the same instance
// when resolved within the same scope.
//
// The scope name should be unique (e.g., a request ID). Creating a scope
// with the same name as an existing scope will replace the old scope.
//
// Example:
//
//	// In an HTTP handler
//	func handler(w http.ResponseWriter, r *http.Request) {
//	    scope := container.CreateScope("request-" + r.Header.Get("X-Request-ID"))
//
//	    // Same instance within this request
//	    ctx1, _ := di.ResolveInScope[*RequestContext](container, scope)
//	    ctx2, _ := di.ResolveInScope[*RequestContext](container, scope)
//	    // ctx1 == ctx2
//	}
func (c *Container) CreateScope(name string) *Scope {
	c.mu.Lock()
	defer c.mu.Unlock()

	scope := newScope(name, c)
	c.scopes[name] = scope
	return scope
}

// ResolveInScope resolves a dependency within a specific scope.
//
// For scoped dependencies (registered with [AsScoped]), the same instance
// is returned for all resolutions within the same scope. Singleton and
// transient dependencies behave normally.
//
// Example:
//
//	di.Register[*RequestContext](c, newRequestContext, di.AsScoped())
//
//	scope := container.CreateScope("request-123")
//	ctx, err := di.ResolveInScope[*RequestContext](container, scope)
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

// Has checks if a type is registered in the container.
//
// This only checks for unnamed registrations. Use [HasNamed] to check
// for named registrations.
//
// Example:
//
//	if di.Has[Logger](container) {
//	    logger := di.MustResolve[Logger](container)
//	    logger.Log("Logger is available")
//	}
func Has[T any](c *Container) bool {
	return HasNamed[T](c, "")
}

// HasNamed checks if a named type is registered in the container.
//
// Example:
//
//	if di.HasNamed[Logger](container, "file") {
//	    fileLogger := di.MustResolveNamed[Logger](container, "file")
//	    // Use file logger
//	}
func HasNamed[T any](c *Container, name string) bool {
	var zero T
	targetType := reflect.TypeOf(&zero).Elem()

	c.mu.RLock()
	defer c.mu.RUnlock()

	key := registrationKey{typ: targetType, name: name}
	_, exists := c.registrations[key]
	return exists
}

// Clear removes all registrations, cached singletons, and scopes from the container.
//
// After calling Clear, the container is empty and new registrations must be made
// before resolving any dependencies.
//
// This is useful in testing scenarios where you want to reset the container
// between tests.
//
// Example:
//
//	func TestSomething(t *testing.T) {
//	    container := di.New()
//	    // ... setup and test ...
//	    container.Clear() // Reset for next test
//	}
func (c *Container) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.registrations = make(map[registrationKey]*registration)
	c.singletons = make(map[registrationKey]any)
	c.scopes = make(map[string]*Scope)
}
