// Package di provides a lightweight, type-safe dependency injection container for Go.
//
// This package leverages Go generics (1.22+) for compile-time type safety while
// providing a clean, intuitive API for managing application dependencies.
//
// # Features
//
//   - Type-safe generics: Compile-time type checking with Register[T] and Resolve[T]
//   - Multiple lifetimes: Transient, Singleton, and Scoped dependency management
//   - Automatic resolution: Constructor parameters are automatically resolved
//   - Circular dependency detection: Fails fast with clear error messages
//   - Named registrations: Register multiple implementations of the same interface
//   - Thread-safe: Safe for concurrent access across goroutines
//   - Zero dependencies: Uses only the Go standard library
//
// # Quick Start
//
// Creating a container and registering dependencies is straightforward:
//
//	container := di.New()
//
//	// Register a singleton logger
//	di.Register[Logger](container, func() Logger {
//	    return &ConsoleLogger{}
//	}, di.AsSingleton())
//
//	// Register a service with auto-resolved dependencies
//	di.Register[UserService](container, func(log Logger) UserService {
//	    return &DefaultUserService{logger: log}
//	})
//
//	// Resolve the service (Logger is automatically injected)
//	service, err := di.Resolve[UserService](container)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// # Dependency Lifetimes
//
// The container supports three dependency lifetimes:
//
// Transient (default): A new instance is created every time the dependency is resolved.
// Use for stateless services or when each consumer needs its own instance.
//
//	di.Register[Service](c, factory)                // default is transient
//	di.Register[Service](c, factory, di.AsTransient())
//
// Singleton: A single instance is created and shared across all resolutions.
// Use for stateless services, configuration objects, or connection pools.
//
//	di.Register[Service](c, factory, di.AsSingleton())
//
// Scoped: A single instance is created per scope. Useful for request-scoped
// dependencies in web applications.
//
//	di.Register[Service](c, factory, di.AsScoped())
//
//	// Later, in a request handler:
//	scope := container.CreateScope("request-123")
//	service, _ := di.ResolveInScope[Service](container, scope)
//
// # Registration Methods
//
// The package provides several ways to register dependencies:
//
// Register with a factory function (most common):
//
//	di.Register[Logger](c, func() Logger {
//	    return &ConsoleLogger{}
//	})
//
// Register with dependencies that are automatically resolved:
//
//	di.Register[UserService](c, func(log Logger, repo UserRepository) UserService {
//	    return &DefaultUserService{logger: log, repo: repo}
//	})
//
// Register a factory that can return an error:
//
//	di.Register[*sql.DB](c, func(cfg Config) (*sql.DB, error) {
//	    return sql.Open("postgres", cfg.DatabaseURL())
//	}, di.AsSingleton())
//
// Register an existing instance:
//
//	config := &AppConfig{Port: 8080}
//	di.RegisterInstance[Config](c, config)
//
// Register an interface to implementation mapping:
//
//	di.RegisterType[UserRepository, PostgresUserRepository](c, di.AsSingleton())
//
// # Named Registrations
//
// Multiple implementations of the same interface can be registered with names:
//
//	di.Register[Logger](c, func() Logger {
//	    return &ConsoleLogger{}
//	}, di.WithName("console"))
//
//	di.Register[Logger](c, func() Logger {
//	    return &FileLogger{path: "/var/log/app.log"}
//	}, di.WithName("file"))
//
//	// Resolve by name
//	consoleLogger, _ := di.ResolveNamed[Logger](c, "console")
//	fileLogger, _ := di.ResolveNamed[Logger](c, "file")
//
// # Error Handling
//
// The package provides typed errors for precise error handling:
//
//	service, err := di.Resolve[UserService](container)
//	if err != nil {
//	    switch e := err.(type) {
//	    case di.ErrNotRegistered:
//	        // Type was not registered
//	    case di.ErrCircularDependency:
//	        // A → B → A dependency chain detected
//	    case di.ErrResolutionFailed:
//	        // Factory returned an error or dependency failed
//	    case di.ErrInvalidFactory:
//	        // Factory signature is invalid
//	    }
//	}
//
// # Thread Safety
//
// The container is fully thread-safe. You can safely register dependencies,
// resolve them, and create scopes from multiple goroutines concurrently.
//
// # Best Practices
//
// 1. Register at startup: Register all dependencies during application startup,
// not at runtime.
//
// 2. Depend on interfaces: Always register and resolve interfaces, not concrete types.
// This makes testing and swapping implementations easier.
//
// 3. Use appropriate lifetimes: Choose Singleton for stateless services and
// shared resources, Transient for stateful objects, and Scoped for per-request data.
//
// 4. Keep factories simple: Factories should only create objects, not perform
// business logic or complex initialization.
package di
