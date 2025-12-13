# Benchmark Results

Performance benchmarks for the Go Dependency Injector.

> **Auto-generated** from CI on $(date -u +"%Y-%m-%d %H:%M UTC")
>
> Runner: GitHub Actions (ubuntu-latest)

## Summary

| Operation | Performance |
|-----------|-------------|
| Container creation | **33.56 ns** |
| Instance resolution | **69.59 ns** |
| Singleton resolution | **89.22 ns** |
| Registration | **107.7 ns** |

---

## Registration Operations

These benchmarks measure the cost of registering dependencies with the container.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkNew` | 33.56 | 0 | 0 |
| `BenchmarkNew` | 33.35 | 0 | 0 |
| `BenchmarkNew` | 35.01 | 0 | 0 |
| `BenchmarkRegister` | 107.7 | 96 | 1 |

---

## Resolution Operations

These benchmarks measure the cost of resolving dependencies from the container.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkResolveTransient` | 309.9 | 56 | 3 |
| `BenchmarkResolveTransient` | 305.3 | 56 | 3 |
| `BenchmarkResolveTransient` | 306.7 | 56 | 3 |
| `BenchmarkResolveSingleton` | 89.22 | 16 | 1 |
| `BenchmarkResolveSingleton` | 89.03 | 16 | 1 |
| `BenchmarkResolveSingleton` | 89.51 | 16 | 1 |

---

## Dependency Chain Resolution

These benchmarks measure resolution performance with dependency injection chains.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkResolveWithOneDependency` | 658.7 | 144 | 7 |
| `BenchmarkResolveWithOneDependency` | 658.0 | 144 | 7 |
| `BenchmarkResolveWithOneDependency` | 652.4 | 144 | 7 |

---

## Parallel/Concurrent Performance

These benchmarks measure thread-safe concurrent resolution.

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkResolveSingletonParallel` | 81.26 | 16 | 1 |
| `BenchmarkResolveSingletonParallel` | 81.23 | 16 | 1 |
| `BenchmarkResolveSingletonParallel` | 81.51 | 16 | 1 |
| `BenchmarkResolveTransientParallel` | 146.6 | 56 | 3 |

---

## Utility Operations

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkHas` | 26.96 | 0 | 0 |
| `BenchmarkHas` | 26.78 | 0 | 0 |
| `BenchmarkHas` | 26.87 | 0 | 0 |

---

## Large Container Performance

| Benchmark | ns/op | B/op | allocs/op |
|-----------|------:|-----:|----------:|
| `BenchmarkContainerWithManyRegistrations` | 10833 | 10208 | 106 |
| `BenchmarkContainerWithManyRegistrations` | 10809 | 10208 | 106 |

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
BenchmarkNew-4                              	32611900	        33.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew-4                              	36025140	        33.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew-4                              	33891902	        35.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkRegister-4                         	11676688	       107.7 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegister-4                         	11734194	       106.8 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegister-4                         	10846932	       106.8 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterWithOptions-4              	10915070	       108.6 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterWithOptions-4              	11116288	       109.2 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterWithOptions-4              	11645382	       106.2 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterInstance-4                 	 9500060	       122.5 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterInstance-4                 	 9804459	       119.0 ns/op	      96 B/op	       1 allocs/op
BenchmarkRegisterInstance-4                 	10509506	       119.2 ns/op	      96 B/op	       1 allocs/op
BenchmarkResolveTransient-4                 	 3922483	       309.9 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransient-4                 	 3932703	       305.3 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransient-4                 	 3925299	       306.7 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveSingleton-4                 	13234936	        89.22 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingleton-4                 	12842667	        89.03 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingleton-4                 	13002880	        89.51 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveInstance-4                  	17267479	        69.59 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveInstance-4                  	17230476	        70.08 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveInstance-4                  	17161243	        69.43 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedSameScope-4           	 9795417	       117.5 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedSameScope-4           	10214576	       116.9 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedSameScope-4           	10131291	       116.2 ns/op	      16 B/op	       1 allocs/op
BenchmarkMustResolve-4                      	12429343	        95.21 ns/op	      16 B/op	       1 allocs/op
BenchmarkMustResolve-4                      	12536749	        96.19 ns/op	      16 B/op	       1 allocs/op
BenchmarkMustResolve-4                      	12442563	        95.83 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveWithOneDependency-4         	 1840736	       658.7 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithOneDependency-4         	 1824277	       658.0 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithOneDependency-4         	 1828780	       652.4 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithTwoDependencies-4       	 1339659	       902.5 ns/op	     232 B/op	       9 allocs/op
BenchmarkResolveWithTwoDependencies-4       	 1323808	       902.4 ns/op	     232 B/op	       9 allocs/op
BenchmarkResolveWithTwoDependencies-4       	 1334764	       897.6 ns/op	     232 B/op	       9 allocs/op
BenchmarkResolveDeepDependencyChain-4       	 1968400	       615.2 ns/op	     128 B/op	       6 allocs/op
BenchmarkResolveDeepDependencyChain-4       	 1961782	       610.4 ns/op	     128 B/op	       6 allocs/op
BenchmarkResolveDeepDependencyChain-4       	 1961894	       612.8 ns/op	     128 B/op	       6 allocs/op
BenchmarkResolveNamed-4                     	13376658	        91.15 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveNamed-4                     	11618572	        89.21 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveNamed-4                     	13234015	        92.01 ns/op	      16 B/op	       1 allocs/op
BenchmarkHas-4                              	45001825	        26.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas-4                              	45028724	        26.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkHas-4                              	45157154	        26.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkHasNamed-4                         	44244534	        27.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkHasNamed-4                         	43797482	        27.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkHasNamed-4                         	43678652	        27.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkCreateScope-4                      	12649666	        96.66 ns/op	     112 B/op	       2 allocs/op
BenchmarkCreateScope-4                      	11596857	        98.51 ns/op	     112 B/op	       2 allocs/op
BenchmarkCreateScope-4                      	11482584	        99.17 ns/op	     112 B/op	       2 allocs/op
BenchmarkResolveSingletonParallel-4         	14675655	        81.26 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingletonParallel-4         	14705733	        81.23 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveSingletonParallel-4         	14772818	        81.51 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveTransientParallel-4         	 8308360	       146.6 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransientParallel-4         	 8273569	       144.8 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveTransientParallel-4         	 8337302	       144.9 ns/op	      56 B/op	       3 allocs/op
BenchmarkResolveScopedParallel-4            	17405810	        79.63 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedParallel-4            	17245372	        69.97 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveScopedParallel-4            	15474723	        78.79 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveWithDepsParallel-4          	 3722482	       318.7 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithDepsParallel-4          	 3724057	       319.6 ns/op	     144 B/op	       7 allocs/op
BenchmarkResolveWithDepsParallel-4          	 3735946	       322.6 ns/op	     144 B/op	       7 allocs/op
BenchmarkContainerWithManyRegistrations-4   	  106766	     10833 ns/op	   10208 B/op	     106 allocs/op
BenchmarkContainerWithManyRegistrations-4   	  104218	     10809 ns/op	   10208 B/op	     106 allocs/op
BenchmarkContainerWithManyRegistrations-4   	  111906	     10761 ns/op	   10208 B/op	     106 allocs/op
BenchmarkResolveFromLargeContainer-4        	13248219	        88.22 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveFromLargeContainer-4        	13037568	        88.71 ns/op	      16 B/op	       1 allocs/op
BenchmarkResolveFromLargeContainer-4        	13338024	        88.76 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/joseph/go-dependency-injection/di	91.983s
```

</details>
