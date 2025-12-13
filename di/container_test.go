package di_test

import (
	"errors"
	"testing"

	"github.com/joseph/go-dependency-injection/di"
)

// =============================================================================
// Test Interfaces and Implementations
// =============================================================================

type Greeter interface {
	Greet(name string) string
}

type SimpleGreeter struct{}

func (g *SimpleGreeter) Greet(name string) string {
	return "Hello, " + name
}

type Logger interface {
	Log(msg string)
}

type TestLogger struct {
	Messages []string
}

func (l *TestLogger) Log(msg string) {
	l.Messages = append(l.Messages, msg)
}

type Service interface {
	DoWork() string
}

type DefaultService struct {
	logger  Logger
	greeter Greeter
}

func (s *DefaultService) DoWork() string {
	s.logger.Log("doing work")
	return s.greeter.Greet("World")
}

// =============================================================================
// Tests
// =============================================================================

func TestRegisterAndResolve(t *testing.T) {
	c := di.New()

	err := di.Register[Greeter](c, func() Greeter {
		return &SimpleGreeter{}
	})
	if err != nil {
		t.Fatalf("failed to register: %v", err)
	}

	greeter, err := di.Resolve[Greeter](c)
	if err != nil {
		t.Fatalf("failed to resolve: %v", err)
	}

	result := greeter.Greet("Test")
	if result != "Hello, Test" {
		t.Errorf("expected 'Hello, Test', got '%s'", result)
	}
}

func TestSingleton(t *testing.T) {
	c := di.New()

	callCount := 0
	err := di.Register[*TestLogger](c, func() *TestLogger {
		callCount++
		return &TestLogger{}
	}, di.AsSingleton())
	if err != nil {
		t.Fatalf("failed to register: %v", err)
	}

	logger1, _ := di.Resolve[*TestLogger](c)
	logger2, _ := di.Resolve[*TestLogger](c)

	if callCount != 1 {
		t.Errorf("factory should be called once for singleton, called %d times", callCount)
	}

	// Verify they're the same instance
	logger1.Log("test")
	if len(logger2.Messages) != 1 {
		t.Error("expected same instance for singleton")
	}
}

func TestTransient(t *testing.T) {
	c := di.New()

	callCount := 0
	err := di.Register[*TestLogger](c, func() *TestLogger {
		callCount++
		return &TestLogger{}
	}, di.AsTransient())
	if err != nil {
		t.Fatalf("failed to register: %v", err)
	}

	_, _ = di.Resolve[*TestLogger](c)
	_, _ = di.Resolve[*TestLogger](c)

	if callCount != 2 {
		t.Errorf("factory should be called twice for transient, called %d times", callCount)
	}
}

func TestRegisterInstance(t *testing.T) {
	c := di.New()

	instance := &TestLogger{Messages: []string{"pre-existing"}}
	di.RegisterInstance[*TestLogger](c, instance)

	resolved, err := di.Resolve[*TestLogger](c)
	if err != nil {
		t.Fatalf("failed to resolve: %v", err)
	}

	if len(resolved.Messages) != 1 || resolved.Messages[0] != "pre-existing" {
		t.Error("expected pre-existing instance")
	}
}

func TestDependencyInjection(t *testing.T) {
	c := di.New()

	// Register Logger
	err := di.Register[Logger](c, func() Logger {
		return &TestLogger{}
	}, di.AsSingleton())
	if err != nil {
		t.Fatalf("failed to register Logger: %v", err)
	}

	// Register Greeter
	err = di.Register[Greeter](c, func() Greeter {
		return &SimpleGreeter{}
	})
	if err != nil {
		t.Fatalf("failed to register Greeter: %v", err)
	}

	// Register Service with dependencies
	err = di.Register[Service](c, func(l Logger, g Greeter) Service {
		return &DefaultService{logger: l, greeter: g}
	})
	if err != nil {
		t.Fatalf("failed to register Service: %v", err)
	}

	// Resolve Service - dependencies should be auto-resolved
	service, err := di.Resolve[Service](c)
	if err != nil {
		t.Fatalf("failed to resolve Service: %v", err)
	}

	result := service.DoWork()
	if result != "Hello, World" {
		t.Errorf("expected 'Hello, World', got '%s'", result)
	}
}

func TestNotRegisteredError(t *testing.T) {
	c := di.New()

	_, err := di.Resolve[Greeter](c)
	if err == nil {
		t.Fatal("expected error for unregistered type")
	}

	var notRegErr di.ErrNotRegistered
	if !errors.As(err, &notRegErr) {
		t.Errorf("expected ErrNotRegistered, got %T", err)
	}
}

func TestCircularDependencyDetection(t *testing.T) {
	c := di.New()

	// A depends on B
	type ServiceA interface{ A() }
	type ServiceB interface{ B() }

	// Create circular dependency: A -> B -> A
	di.Register[ServiceA](c, func(b ServiceB) ServiceA {
		return nil // Implementation doesn't matter for this test
	})
	di.Register[ServiceB](c, func(a ServiceA) ServiceB {
		return nil
	})

	_, err := di.Resolve[ServiceA](c)
	if err == nil {
		t.Fatal("expected circular dependency error")
	}

	var circErr di.ErrCircularDependency
	if !errors.As(err, &circErr) {
		t.Errorf("expected ErrCircularDependency, got %T: %v", err, err)
	}
}

func TestNamedRegistrations(t *testing.T) {
	c := di.New()

	// Register two different Greeter implementations with names
	err := di.Register[Greeter](c, func() Greeter {
		return &SimpleGreeter{}
	}, di.WithName("simple"))
	if err != nil {
		t.Fatalf("failed to register simple: %v", err)
	}

	err = di.Register[Greeter](c, func() Greeter {
		return &formalGreeter{}
	}, di.WithName("formal"))
	if err != nil {
		t.Fatalf("failed to register formal: %v", err)
	}

	simple, err := di.ResolveNamed[Greeter](c, "simple")
	if err != nil {
		t.Fatalf("failed to resolve simple: %v", err)
	}

	formal, err := di.ResolveNamed[Greeter](c, "formal")
	if err != nil {
		t.Fatalf("failed to resolve formal: %v", err)
	}

	if simple.Greet("Test") != "Hello, Test" {
		t.Error("simple greeter incorrect")
	}

	if formal.Greet("Test") != "Good day, Test" {
		t.Error("formal greeter incorrect")
	}
}

type formalGreeter struct{}

func (g *formalGreeter) Greet(name string) string {
	return "Good day, " + name
}

func TestHas(t *testing.T) {
	c := di.New()

	if di.Has[Greeter](c) {
		t.Error("expected Has to return false before registration")
	}

	di.Register[Greeter](c, func() Greeter { return &SimpleGreeter{} })

	if !di.Has[Greeter](c) {
		t.Error("expected Has to return true after registration")
	}
}

func TestMustResolve(t *testing.T) {
	c := di.New()
	di.Register[Greeter](c, func() Greeter { return &SimpleGreeter{} })

	// Should not panic
	greeter := di.MustResolve[Greeter](c)
	if greeter == nil {
		t.Error("expected greeter instance")
	}
}

func TestMustResolvePanics(t *testing.T) {
	c := di.New()

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for unregistered type")
		}
	}()

	di.MustResolve[Greeter](c)
}

func TestFactoryWithErrorReturn(t *testing.T) {
	c := di.New()

	expectedErr := errors.New("factory error")

	err := di.Register[Greeter](c, func() (Greeter, error) {
		return nil, expectedErr
	})
	if err != nil {
		t.Fatalf("failed to register: %v", err)
	}

	_, err = di.Resolve[Greeter](c)
	if err == nil {
		t.Fatal("expected error from factory")
	}

	var resErr di.ErrResolutionFailed
	if !errors.As(err, &resErr) {
		t.Errorf("expected ErrResolutionFailed, got %T", err)
	}
}

func TestScopedLifetime(t *testing.T) {
	c := di.New()

	callCount := 0
	err := di.Register[*TestLogger](c, func() *TestLogger {
		callCount++
		return &TestLogger{}
	}, di.AsScoped())
	if err != nil {
		t.Fatalf("failed to register: %v", err)
	}

	// Create two scopes
	scope1 := c.CreateScope("scope1")
	scope2 := c.CreateScope("scope2")

	// Resolve in scope1 twice - should return same instance
	logger1a, _ := di.ResolveInScope[*TestLogger](c, scope1)
	logger1b, _ := di.ResolveInScope[*TestLogger](c, scope1)

	// Resolve in scope2 - should return different instance
	logger2, _ := di.ResolveInScope[*TestLogger](c, scope2)

	// Verify scope1 returns same instance
	logger1a.Log("from 1a")
	if len(logger1b.Messages) != 1 {
		t.Error("expected same instance within scope")
	}

	// Verify scope2 has different instance
	if len(logger2.Messages) != 0 {
		t.Error("expected different instance in different scope")
	}

	// Should have created 2 instances (one per scope)
	if callCount != 2 {
		t.Errorf("expected 2 factory calls for 2 scopes, got %d", callCount)
	}
}

func TestClear(t *testing.T) {
	c := di.New()

	di.Register[Greeter](c, func() Greeter { return &SimpleGreeter{} })

	if !di.Has[Greeter](c) {
		t.Error("expected registration to exist")
	}

	c.Clear()

	if di.Has[Greeter](c) {
		t.Error("expected registration to be cleared")
	}
}

func TestInvalidFactory(t *testing.T) {
	c := di.New()

	// Not a function
	err := di.Register[Greeter](c, "not a function")
	if err == nil {
		t.Fatal("expected error for non-function factory")
	}

	var invalidErr di.ErrInvalidFactory
	if !errors.As(err, &invalidErr) {
		t.Errorf("expected ErrInvalidFactory, got %T", err)
	}
}

func TestInvalidFactoryNoReturn(t *testing.T) {
	c := di.New()

	err := di.Register[Greeter](c, func() {})
	if err == nil {
		t.Fatal("expected error for factory with no return")
	}
}

func TestInvalidFactoryWrongReturnType(t *testing.T) {
	c := di.New()

	err := di.Register[Greeter](c, func() string { return "wrong" })
	if err == nil {
		t.Fatal("expected error for factory with wrong return type")
	}
}

