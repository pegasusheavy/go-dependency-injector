# Benchmark Results

Performance benchmarks for the Go Dependency Injector.

> **Auto-generated** from CI on $(date -u +"%Y-%m-%d %H:%M UTC")
>
> Runner: GitHub Actions (ubuntu-latest)

## Summary

| Operation | Performance |
|-----------|-------------|
| Container creation | **33.77 ns** |
| Instance resolution | **69.06 ns** |
| Singleton resolution | **88.19 ns** |
| Registration | **100.9 ns** |

---

## Registration Operations

These benchmarks measure the cost of registering dependencies with the container.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkNew` | 33.77 | 0 | 0 |
| `BenchmarkNew` | 33.39 | 0 | 0 |
| `BenchmarkNew` | 33.39 | 0 | 0 |
| `BenchmarkRegister` | 100.9 | 96 | 1 |

---

## Resolution Operations

These benchmarks measure the cost of resolving dependencies from the container.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkResolveTransient` | 308.2 | 56 | 3 |
| `BenchmarkResolveTransient` | 303.1 | 56 | 3 |
| `BenchmarkResolveTransient` | 303.7 | 56 | 3 |
| `BenchmarkResolveSingleton` | 88.19 | 16 | 1 |
| `BenchmarkResolveSingleton` | 88.49 | 16 | 1 |
| `BenchmarkResolveSingleton` | 88.45 | 16 | 1 |

---

## Dependency Chain Resolution

These benchmarks measure resolution performance with dependency injection chains.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkResolveWithOneDependency` | 648.1 | 144 | 7 |
| `BenchmarkResolveWithOneDependency` | 648.0 | 144 | 7 |
| `BenchmarkResolveWithOneDependency` | 648.6 | 144 | 7 |

---

## Parallel/Concurrent Performance

These benchmarks measure thread-safe concurrent resolution.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkResolveSingletonParallel` | 79.51 | 16 | 1 |
| `BenchmarkResolveSingletonParallel` | 79.51 | 16 | 1 |
| `BenchmarkResolveSingletonParallel` | 79.43 | 16 | 1 |
| `BenchmarkResolveTransientParallel` | 146.3 | 56 | 3 |

---

## Utility Operations

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkHas` | 26.98 | 0 | 0 |
| `BenchmarkHas` | 26.84 | 0 | 0 |
| `BenchmarkHas` | 26.97 | 0 | 0 |

---

## Large Container Performance

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkContainerWithManyRegistrations` | 10473 | 10208 | 106 |
| `BenchmarkContainerWithManyRegistrations` | 10648 | 10208 | 106 |

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
BenchmarkNew-4                              	35281148	        33.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew-4                              	35783170	        33.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew-4                              	35584957	        33.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkRegister-4                         	11784241	       100.9 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegister-4                         	11961526	       101.4 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegister-4                         	11558036	       104.7 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterWithOptions-4              	11021912	       104.8 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterWithOptions-4              	11429470	       103.8 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterWithOptions-4              	11639494	       106.5 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterInstance-4                 	10552128	       114.2 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterInstance-4                 	 9849888	       115.1 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterInstance-4                 	10388235	       115.2 ns/op	      96 B/op	       1 allocs/op
BenchmarkResolveTransient-4                 	 3933405	       308.2 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransient-4                 	 3914638	       303.1 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransient-4                 	 3931684	       303.7 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveSingleton-4                 	13293129	        88.19 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingleton-4                 	13172013	        88.49 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingleton-4                 	13297744	        88.45 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveInstance-4                  	17291788	        69.06 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveInstance-4                  	16988540	        69.02 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveInstance-4                  	17026098	        69.97 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedSameScope-4           	 9716161	       115.8 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedSameScope-4           	10226342	       116.6 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedSameScope-4           	10220260	       116.2 ns/op	      16 B/op	       1 allocs/op
BenchmarkMustResolve-4                      	12339110	        95.12 ns/op	      16 B/op	       1 allocs/op
BenchmarkMustResolve-4                      	12348453	        95.30 ns/op	      16 B/op	       1 allocs/op
BenchmarkMustResolve-4                      	12260962	        95.11 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveWithOneDependency-4         	 1847868	       648.1 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithOneDependency-4         	 1826055	       648.0 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithOneDependency-4         	 1840140	       648.6 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithTwoDependencies-4       	 1356070	       885.8 ns/op	     232 B/op	       9 allocs/op
BenchmarkResolveWithTwoDependencies-4       	 1343416	       893.8 ns/op	     232 B/op	       9 allocs/op
BenchmarkResolveWithTwoDependencies-4       	 1348425	       897.8 ns/op	     232 B/op	       9 allocs/op
BenchmarkResolveDeepDependencyChain-4       	 1965902	       608.1 ns/op	     128 B/op	       6 allocs/op
BenchmarkResolveDeepDependencyChain-4       	 1952476	       609.1 ns/op	     128 B/op	       6 allocs/op
BenchmarkResolveDeepDependencyChain-4       	 1967620	       607.6 ns/op	     128 B/op	       6 allocs/op
BenchmarkResolveNamed-4                     	13226448	        88.39 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveNamed-4                     	13153689	        88.80 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveNamed-4                     	13282407	        88.49 ns/op	      16 B/op	       1 allocs/op
BenchmarkHas-4                              	45095649	        26.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas-4                              	45008785	        26.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas-4                              	45273504	        26.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkHasNamed-4                         	44157086	        27.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkHasNamed-4                         	44165226	        27.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkHasNamed-4                         	40480796	        27.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkCreateScope-4                      	12137769	        94.70 ns/op	     112 B/op	       2 allocs/op
BenchmarkCreateScope-4                      	12512762	        95.69 ns/op	     112 B/op	       2 allocs/op
BenchmarkCreateScope-4                      	12355357	        97.44 ns/op	     112 B/op	       2 allocs/op
BenchmarkResolveSingletonParallel-4         	14875749	        79.51 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingletonParallel-4         	14906710	        79.51 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingletonParallel-4         	14697283	        79.43 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveTransientParallel-4         	 8238846	       146.3 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransientParallel-4         	 8212813	       146.3 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransientParallel-4         	 8215731	       145.9 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveScopedParallel-4            	17227365	        66.84 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedParallel-4            	17765860	        68.74 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedParallel-4            	17491858	        66.88 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveWithDepsParallel-4          	 3675274	       323.6 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithDepsParallel-4          	 3690685	       322.1 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithDepsParallel-4          	 3726633	       322.4 ns/op	     144 B/op	       7 allocs/op
BenchmarkContainerWithManyRegistrations-4   	  115585	     10473 ns/op	   10208 B/op	     106 allocs/op
BenchmarkContainerWithManyRegistrations-4   	  113146	     10648 ns/op	   10208 B/op	     106 allocs/op
BenchmarkContainerWithManyRegistrations-4   	  104773	     10810 ns/op	   10208 B/op	     106 allocs/op
BenchmarkResolveFromLargeContainer-4        	13298329	        88.30 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveFromLargeContainer-4        	13277791	        87.98 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveFromLargeContainer-4        	13366665	        88.19 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/pegasusheavy/go-dependency-injector/di	91.399s
```

</details>
