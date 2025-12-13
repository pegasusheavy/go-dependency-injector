# Benchmark Results

Performance benchmarks for the Go Dependency Injector.

> **Auto-generated** from CI on $(date -u +"%Y-%m-%d %H:%M UTC")
>
> Runner: GitHub Actions (ubuntu-latest)

## Summary

| Operation | Performance |
|-----------|-------------|
| Container creation | **34.73 ns** |
| Instance resolution | **70.78 ns** |
| Singleton resolution | **88.93 ns** |
| Registration | **106.2 ns** |

---

## Registration Operations

These benchmarks measure the cost of registering dependencies with the container.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkNew` | 34.73 | 0 | 0 |
| `BenchmarkNew` | 34.71 | 0 | 0 |
| `BenchmarkNew` | 33.88 | 0 | 0 |
| `BenchmarkRegister` | 106.2 | 96 | 1 |

---

## Resolution Operations

These benchmarks measure the cost of resolving dependencies from the container.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkResolveTransient` | 303.0 | 56 | 3 |
| `BenchmarkResolveTransient` | 304.1 | 56 | 3 |
| `BenchmarkResolveTransient` | 303.1 | 56 | 3 |
| `BenchmarkResolveSingleton` | 88.93 | 16 | 1 |
| `BenchmarkResolveSingleton` | 87.94 | 16 | 1 |
| `BenchmarkResolveSingleton` | 88.30 | 16 | 1 |

---

## Dependency Chain Resolution

These benchmarks measure resolution performance with dependency injection chains.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkResolveWithOneDependency` | 648.2 | 144 | 7 |
| `BenchmarkResolveWithOneDependency` | 647.7 | 144 | 7 |
| `BenchmarkResolveWithOneDependency` | 647.8 | 144 | 7 |

---

## Parallel/Concurrent Performance

These benchmarks measure thread-safe concurrent resolution.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkResolveSingletonParallel` | 79.96 | 16 | 1 |
| `BenchmarkResolveSingletonParallel` | 80.37 | 16 | 1 |
| `BenchmarkResolveSingletonParallel` | 80.09 | 16 | 1 |
| `BenchmarkResolveTransientParallel` | 143.4 | 56 | 3 |

---

## Utility Operations

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkHas` | 27.23 | 0 | 0 |
| `BenchmarkHas` | 26.79 | 0 | 0 |
| `BenchmarkHas` | 27.06 | 0 | 0 |

---

## Large Container Performance

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkContainerWithManyRegistrations` | 10927 | 10208 | 106 |
| `BenchmarkContainerWithManyRegistrations` | 10851 | 10208 | 106 |

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
pkg: github.com/joseph/go-dependency-injection/di
cpu: AMD EPYC 7763 64-Core Processor                
BenchmarkNew-4                              	33776436	        34.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew-4                              	35762455	        34.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew-4                              	35747758	        33.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkRegister-4                         	11433500	       106.2 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegister-4                         	11317702	       106.1 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegister-4                         	10832203	       102.0 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterWithOptions-4              	10723082	       106.0 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterWithOptions-4              	11350717	       108.3 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterWithOptions-4              	10868216	       108.5 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterInstance-4                 	10223548	       121.4 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterInstance-4                 	 9314716	       118.6 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterInstance-4                 	10502535	       117.8 ns/op	      96 B/op	       1 allocs/op
BenchmarkResolveTransient-4                 	 3948649	       303.0 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransient-4                 	 3938739	       304.1 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransient-4                 	 3906367	       303.1 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveSingleton-4                 	13524258	        88.93 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingleton-4                 	13283160	        87.94 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingleton-4                 	13372344	        88.30 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveInstance-4                  	17137120	        70.78 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveInstance-4                  	17170867	        68.99 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveInstance-4                  	17219488	        68.73 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedSameScope-4           	10269906	       116.2 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedSameScope-4           	10234885	       116.7 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedSameScope-4           	 9929479	       116.2 ns/op	      16 B/op	       1 allocs/op
BenchmarkMustResolve-4                      	12401398	        95.17 ns/op	      16 B/op	       1 allocs/op
BenchmarkMustResolve-4                      	12384904	        95.71 ns/op	      16 B/op	       1 allocs/op
BenchmarkMustResolve-4                      	12378092	        95.10 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveWithOneDependency-4         	 1857562	       648.2 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithOneDependency-4         	 1847209	       647.7 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithOneDependency-4         	 1841114	       647.8 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithTwoDependencies-4       	 1359033	       882.8 ns/op	     232 B/op	       9 allocs/op
BenchmarkResolveWithTwoDependencies-4       	 1356098	       893.5 ns/op	     232 B/op	       9 allocs/op
BenchmarkResolveWithTwoDependencies-4       	 1352257	       884.4 ns/op	     232 B/op	       9 allocs/op
BenchmarkResolveDeepDependencyChain-4       	 1951630	       604.5 ns/op	     128 B/op	       6 allocs/op
BenchmarkResolveDeepDependencyChain-4       	 1964244	       605.5 ns/op	     128 B/op	       6 allocs/op
BenchmarkResolveDeepDependencyChain-4       	 1978966	       606.2 ns/op	     128 B/op	       6 allocs/op
BenchmarkResolveNamed-4                     	13087311	        89.32 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveNamed-4                     	13390800	        88.22 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveNamed-4                     	13318485	        89.09 ns/op	      16 B/op	       1 allocs/op
BenchmarkHas-4                              	45304934	        27.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas-4                              	45098341	        26.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas-4                              	45090378	        27.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkHasNamed-4                         	44195832	        27.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkHasNamed-4                         	44197311	        27.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkHasNamed-4                         	44148716	        27.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkCreateScope-4                      	11458155	        99.03 ns/op	     112 B/op	       2 allocs/op
BenchmarkCreateScope-4                      	12163464	        97.79 ns/op	     112 B/op	       2 allocs/op
BenchmarkCreateScope-4                      	11899947	        98.50 ns/op	     112 B/op	       2 allocs/op
BenchmarkResolveSingletonParallel-4         	14942358	        79.96 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingletonParallel-4         	14937466	        80.37 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingletonParallel-4         	14937096	        80.09 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveTransientParallel-4         	 8126022	       143.4 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransientParallel-4         	 8385739	       144.2 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransientParallel-4         	 8364154	       143.9 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveScopedParallel-4            	17596152	        67.58 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedParallel-4            	17656528	        67.96 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedParallel-4            	17492818	        67.65 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveWithDepsParallel-4          	 3754281	       325.7 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithDepsParallel-4          	 3768733	       318.2 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithDepsParallel-4          	 3769582	       318.3 ns/op	     144 B/op	       7 allocs/op
BenchmarkContainerWithManyRegistrations-4   	  102501	     10927 ns/op	   10208 B/op	     106 allocs/op
BenchmarkContainerWithManyRegistrations-4   	  110564	     10851 ns/op	   10208 B/op	     106 allocs/op
BenchmarkContainerWithManyRegistrations-4   	  108031	     10731 ns/op	   10208 B/op	     106 allocs/op
BenchmarkResolveFromLargeContainer-4        	13370712	        87.99 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveFromLargeContainer-4        	13109805	        88.28 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveFromLargeContainer-4        	13346176	        88.12 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/joseph/go-dependency-injection/di	91.570s
```

</details>
