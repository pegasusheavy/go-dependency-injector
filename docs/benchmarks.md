# Benchmark Results

Performance benchmarks for the Go Dependency Injector, measured on an Intel Core i9-14900HX.

## Summary

| Operation | Performance |
|-----------|-------------|
| Container creation | **34 ns** |
| Singleton resolution | **67-100 ns** |
| Transient resolution | **303 ns** |
| Registration | **100 ns** |

---

## Registration Operations

These benchmarks measure the cost of registering dependencies with the container.

| Benchmark | ns/op | B/op | allocs/op | Description |
|-----------|------:|-----:|----------:|-------------|
| `New` | 34.43 | 0 | 0 | Create empty container |
| `Register` | 100.2 | 96 | 1 | Register with factory |
| `RegisterWithOptions` | 100.2 | 96 | 1 | Register with lifetime option |
| `RegisterInstance` | 110.8 | 96 | 1 | Register pre-created instance |

### Key Insights

- Creating a new container is extremely fast (~34ns) with zero allocations
- Registration operations are consistent at ~100ns regardless of options
- Memory footprint per registration is minimal (96 bytes)

---

## Resolution Operations

These benchmarks measure the cost of resolving dependencies from the container.

| Benchmark | ns/op | B/op | allocs/op | Description |
|-----------|------:|-----:|----------:|-------------|
| `ResolveInstance` | 66.81 | 16 | 1 | Resolve pre-registered instance |
| `MustResolve` | 86.50 | 16 | 1 | Resolve with panic on error |
| `ResolveNamed` | 86.81 | 16 | 1 | Resolve named registration |
| `ResolveScopedSameScope` | 95.09 | 16 | 1 | Resolve scoped (cached) |
| `ResolveSingleton` | 100.5 | 16 | 1 | Resolve singleton (cached) |
| `ResolveTransient` | 303.4 | 56 | 3 | Resolve transient (new instance) |

### Key Insights

- **Instance resolution is fastest** (~67ns) - no factory invocation needed
- **Singletons are 3x faster** than transients after initial creation
- **Scoped resolution** performs similarly to singletons when cached
- Transient resolution requires more allocations due to factory invocation

---

## Dependency Chain Resolution

These benchmarks measure resolution performance with dependency injection chains.

| Benchmark | ns/op | B/op | allocs/op | Description |
|-----------|------:|-----:|----------:|-------------|
| `ResolveWithOneDependency` | 625.6 | 144 | 7 | Service → Logger |
| `ResolveDeepDependencyChain` | 578.3 | 128 | 6 | 5-level deep chain (singletons) |
| `ResolveWithTwoDependencies` | 951.1 | 232 | 9 | Service → Logger + Service |

### Key Insights

- Deep singleton chains can be **faster** than shallow transient chains
- Each dependency level adds ~100-200ns overhead
- Using singletons for shared dependencies significantly improves performance

---

## Parallel/Concurrent Performance

These benchmarks measure thread-safe concurrent resolution.

| Benchmark | ns/op | B/op | allocs/op | Description |
|-----------|------:|-----:|----------:|-------------|
| `ResolveScopedParallel` | 135.7 | 16 | 1 | Concurrent scoped resolution |
| `ResolveTransientParallel` | 165.9 | 56 | 3 | Concurrent transient resolution |
| `ResolveSingletonParallel` | 227.5 | 16 | 1 | Concurrent singleton resolution |
| `ResolveWithDepsParallel` | 446.9 | 144 | 7 | Concurrent with dependencies |

### Key Insights

- All operations are **thread-safe** with minimal lock contention
- Scoped resolution has the best parallel performance
- Lock overhead is ~100-150ns for singleton resolution

---

## Utility Operations

| Benchmark | ns/op | B/op | allocs/op | Description |
|-----------|------:|-----:|----------:|-------------|
| `Has` | 30.14 | 0 | 0 | Check if type is registered |
| `HasNamed` | 29.46 | 0 | 0 | Check named registration |
| `CreateScope` | 114.6 | 112 | 2 | Create new resolution scope |

### Key Insights

- `Has` checks are extremely fast (~30ns) with zero allocations
- Scope creation is lightweight (~115ns)

---

## Large Container Performance

| Benchmark | ns/op | B/op | allocs/op | Description |
|-----------|------:|-----:|----------:|-------------|
| `ContainerWithManyRegistrations` | 11,115 | 10,208 | 106 | Create + 100 registrations |
| `ResolveFromLargeContainer` | 80.64 | 16 | 1 | Resolve from 100-registration container |

### Key Insights

- Resolution time is **constant** regardless of container size
- Large containers scale well with O(1) lookup performance

---

## Recommendations

Based on these benchmarks:

1. **Use Singletons** for stateless services - 3x faster than transients
2. **Use RegisterInstance** for pre-created objects - fastest resolution
3. **Use Scoped** for request-scoped dependencies - best parallel performance
4. **Avoid deep transient chains** - each level adds allocation overhead

---

## Running Benchmarks

To run these benchmarks yourself:

```bash
# Run all benchmarks
go test ./di/... -bench=. -benchmem

# Run specific benchmark
go test ./di/... -bench=BenchmarkResolveSingleton -benchmem

# Run with more iterations for accuracy
go test ./di/... -bench=. -benchmem -count=5
```

---

*Benchmarks run on Linux (WSL2) with Go 1.22+ on Intel Core i9-14900HX*

