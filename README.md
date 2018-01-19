# goroutine-pool
A simple goroutine pool which can create and release goroutine dynamically, inspired by fasthttp.

# benchmarks
```
Running tool: D:\Go\bin\go.exe test -benchmem -run=^$ workerpool -bench ^BenchmarkGoroutine$

goos: windows
goarch: amd64
pkg: workerpool
BenchmarkGoroutine-4   	    1000	   1634621 ns/op	      65 B/op	       1 allocs/op
PASS
ok  	workerpool	2.047s
Success: Benchmarks passed.
```

```
goos: windows
goarch: amd64
pkg: workerpool
BenchmarkPool-4   	    2000	   1198307 ns/op	      17 B/op	       1 allocs/op
PASS
ok  	workerpool	2.835s
Success: Benchmarks passed.
```
