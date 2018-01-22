# goroutine-pool
[![Build Status](https://travis-ci.org/hnlq715/goroutine-pool.svg?branch=master)](https://travis-ci.org/hnlq715/goroutine-pool)

A simple goroutine pool which can create and release goroutine dynamically, inspired by fasthttp.

# install
```
go get -u -v github.com/hnlq715/goroutine-pool
```
# example
```
package main

import (
	"fmt"
	pool "github.com/hnlq715/goroutine-pool"
	"sync"
)

func main() {
	pool.Start()
	defer pool.Stop()

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		pool.Go(func() {
			fmt.Println(i)
			wg.Done()
		})
	}
	wg.Wait()
	fmt.Println()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		n := i
		pool.Go(func() {
			fmt.Println(n)
			wg.Done()
		})
	}
	wg.Wait()
}
```

```
$ go run main.go
5
5
5
5
5

4
1
0
2
3
```

# benchmarks

## With Wait
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
Running tool: D:\Go\bin\go.exe test -benchmem -run=^$ workerpool -bench ^BenchmarkPool$

goos: windows
goarch: amd64
pkg: workerpool
BenchmarkPool-4   	    2000	   1146818 ns/op	      17 B/op	       1 allocs/op
PASS
ok  	workerpool	2.702s
Success: Benchmarks passed.
```

## Without Wait

### cpu run at 100%
```
Running tool: D:\Go\bin\go.exe test -benchmem -run=^$ workerpool -bench ^BenchmarkGoroutineWithoutWait$

goos: windows
goarch: amd64
pkg: workerpool
BenchmarkGoroutineWithoutWait-4   	 1000000	      4556 ns/op	     517 B/op	       1 allocs/op
PASS
ok  	workerpool	5.649s
Success: Benchmarks passed.
```

### cpu relatively low

```
Running tool: D:\Go\bin\go.exe test -benchmem -run=^$ workerpool -bench ^BenchmarkPoolWithoutWait$

goos: windows
goarch: amd64
pkg: workerpool
BenchmarkPoolWithoutWait-4   	10000000	       144 ns/op	       3 B/op	       0 allocs/op
PASS
ok  	workerpool	4.812s
Success: Benchmarks passed.
```
