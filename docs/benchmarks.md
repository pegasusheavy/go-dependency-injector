# Benchmark Results

Performance benchmarks for the Go Dependency Injector.

> **Auto-generated** from CI on $(date -u +"%Y-%m-%d %H:%M UTC")
>
> Runner: GitHub Actions (ubuntu-latest)

## Summary

| Operation | Performance |
|-----------|-------------|
| Container creation | **33.41 ns** |
| Instance resolution | **69.97 ns** |
| Singleton resolution | **89.45 ns** |
| Registration | **109.6 ns** |

---

## Registration Operations

These benchmarks measure the cost of registering dependencies with the container.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkNew` | 33.41 | 0 | 0 |
| `BenchmarkNew` | 33.53 | 0 | 0 |
| `BenchmarkNew` | 33.44 | 0 | 0 |
| `BenchmarkRegister` | 109.6 | 96 | 1 |

---

## Resolution Operations

These benchmarks measure the cost of resolving dependencies from the container.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkResolveTransient` | 305.5 | 56 | 3 |
| `BenchmarkResolveTransient` | 306.4 | 56 | 3 |
| `BenchmarkResolveTransient` | 311.4 | 56 | 3 |
| `BenchmarkResolveSingleton` | 89.45 | 16 | 1 |
| `BenchmarkResolveSingleton` | 88.32 | 16 | 1 |
| `BenchmarkResolveSingleton` | 90.18 | 16 | 1 |

---

## Dependency Chain Resolution

These benchmarks measure resolution performance with dependency injection chains.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkResolveWithOneDependency` | 645.7 | 144 | 7 |
| `BenchmarkResolveWithOneDependency` | 650.9 | 144 | 7 |
| `BenchmarkResolveWithOneDependency` | 651.0 | 144 | 7 |

---

## Parallel/Concurrent Performance

These benchmarks measure thread-safe concurrent resolution.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkResolveSingletonParallel` | 79.24 | 16 | 1 |
| `BenchmarkResolveSingletonParallel` | 80.33 | 16 | 1 |
| `BenchmarkResolveSingletonParallel` | 79.53 | 16 | 1 |
| `BenchmarkResolveTransientParallel` | 144.8 | 56 | 3 |

---

## Utility Operations

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkHas` | 26.78 | 0 | 0 |
| `BenchmarkHas` | 26.80 | 0 | 0 |
| `BenchmarkHas` | 26.80 | 0 | 0 |

---

## Large Container Performance

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkContainerWithManyRegistrations` | 10847 | 10208 | 106 |
| `BenchmarkContainerWithManyRegistrations` | 10800 | 10208 | 106 |

---

## Recommendations

Based on these benchmarks:

1. **Use Singletons** for stateless services — significantly faster than transients after initial creation
2. **Use RegisterInstance** for pre-created objects — fastest resolution path
3. **Use Scoped** for request-scoped dependencies — excellent parallel performance
4. **Minimize transient chains** — each level adds allocation overhead

---

## Running Benchmarks Locally

```bash
# Run all benchmarks
go test ./di/... -bench=. -benchmem

# Run specific benchmark
go test ./di/... -bench=BenchmarkResolveSingleton -benchmem

# Run with multiple iterations for accuracy
go test ./di/... -bench=. -benchmem -count=5
```

---

## Raw Output

<details>
<summary>Click to expand full benchmark output</summary>

```
goos: linux
goarch: amd64
pkg: github.com/pegasusheavy/go-dependency-injector/di
cpu: AMD EPYC 7763 64-Core Processor                
BenchmarkNew-4                              	35229590	        33.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew-4                              	35927958	        33.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew-4                              	35032476	        33.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkRegister-4                         	11942306	       109.6 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegister-4                         	11114854	       104.2 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegister-4                         	11364734	       103.1 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterWithOptions-4              	10331648	       107.2 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterWithOptions-4              	11415102	       103.0 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterWithOptions-4              	11739834	       102.8 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterInstance-4                 	10527387	       112.6 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterInstance-4                 	10650298	       113.5 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterInstance-4                 	 9951987	       119.0 ns/op	      96 B/op	       1 allocs/op
BenchmarkResolveTransient-4                 	 3836755	       305.5 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransient-4                 	 3949737	       306.4 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransient-4                 	 3875955	       311.4 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveSingleton-4                 	13296450	        89.45 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingleton-4                 	13163416	        88.32 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingleton-4                 	13312137	        90.18 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveInstance-4                  	16714744	        69.97 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveInstance-4                  	17160535	        70.14 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveInstance-4                  	16496871	        70.68 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedSameScope-4           	10228382	       117.3 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedSameScope-4           	10129645	       117.6 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedSameScope-4           	10151343	       117.0 ns/op	      16 B/op	       1 allocs/op
BenchmarkMustResolve-4                      	12223982	        95.62 ns/op	      16 B/op	       1 allocs/op
BenchmarkMustResolve-4                      	12441210	        95.73 ns/op	      16 B/op	       1 allocs/op
BenchmarkMustResolve-4                      	12396130	        96.51 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveWithOneDependency-4         	 1822923	       645.7 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithOneDependency-4         	 1839364	       650.9 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithOneDependency-4         	 1833801	       651.0 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithTwoDependencies-4       	 1343624	       891.3 ns/op	     232 B/op	       9 allocs/op
BenchmarkResolveWithTwoDependencies-4       	 1338019	       887.2 ns/op	     232 B/op	       9 allocs/op
BenchmarkResolveWithTwoDependencies-4       	 1344846	       889.0 ns/op	     232 B/op	       9 allocs/op
BenchmarkResolveDeepDependencyChain-4       	 1979361	       612.7 ns/op	     128 B/op	       6 allocs/op
BenchmarkResolveDeepDependencyChain-4       	 1982600	       619.5 ns/op	     128 B/op	       6 allocs/op
BenchmarkResolveDeepDependencyChain-4       	 1957856	       614.6 ns/op	     128 B/op	       6 allocs/op
BenchmarkResolveNamed-4                     	13238864	        90.26 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveNamed-4                     	13111846	        88.07 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveNamed-4                     	12757592	        88.21 ns/op	      16 B/op	       1 allocs/op
BenchmarkHas-4                              	44919006	        26.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas-4                              	45560803	        26.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas-4                              	45287250	        26.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkHasNamed-4                         	44128734	        27.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkHasNamed-4                         	44052150	        27.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkHasNamed-4                         	43455164	        27.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkCreateScope-4                      	11879060	        98.61 ns/op	     112 B/op	       2 allocs/op
BenchmarkCreateScope-4                      	11783727	        98.70 ns/op	     112 B/op	       2 allocs/op
BenchmarkCreateScope-4                      	11767101	        97.63 ns/op	     112 B/op	       2 allocs/op
BenchmarkResolveSingletonParallel-4         	15072459	        79.24 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingletonParallel-4         	15031858	        80.33 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingletonParallel-4         	15100272	        79.53 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveTransientParallel-4         	 8255784	       144.8 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransientParallel-4         	 8262848	       145.1 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransientParallel-4         	 8252396	       145.4 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveScopedParallel-4            	17471204	        75.88 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedParallel-4            	15773002	        76.63 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedParallel-4            	17943892	        78.04 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveWithDepsParallel-4          	 3765463	       319.4 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithDepsParallel-4          	 3748669	       318.6 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithDepsParallel-4          	 3752712	       318.8 ns/op	     144 B/op	       7 allocs/op
BenchmarkContainerWithManyRegistrations-4   	  103826	     10847 ns/op	   10208 B/op	     106 allocs/op
BenchmarkContainerWithManyRegistrations-4   	  111187	     10800 ns/op	   10208 B/op	     106 allocs/op
BenchmarkContainerWithManyRegistrations-4   	  113468	     10713 ns/op	   10208 B/op	     106 allocs/op
BenchmarkResolveFromLargeContainer-4        	13398297	        88.88 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveFromLargeContainer-4        	12788198	        88.97 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveFromLargeContainer-4        	13290085	        88.37 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/pegasusheavy/go-dependency-injector/di	91.945s
```

</details>
