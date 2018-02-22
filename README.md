# prand

## Benchmark

```
goos: darwin
goarch: amd64
pkg: github.com/mono0x/prand
BenchmarkMathRandInt63-8                100000000               22.2 ns/op       361.02 MB/s
BenchmarkPrandInt63-8                   20000000                94.9 ns/op        84.27 MB/s
BenchmarkMathRandInt63Parallel-8        20000000               103 ns/op          77.62 MB/s
BenchmarkPrandInt63Parallel-8           50000000                24.7 ns/op       323.91 MB/s
BenchmarkMathRandUint64-8               100000000               22.8 ns/op       350.91 MB/s
BenchmarkPrandUint64-8                  20000000               105 ns/op          75.77 MB/s
BenchmarkMathRandUint64Parallel-8       20000000               104 ns/op          76.75 MB/s
BenchmarkPrandUint64Parallel-8          50000000                26.6 ns/op       301.27 MB/s
PASS
ok      github.com/mono0x/prand 15.787s
```
