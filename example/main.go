package main

import (
	"fmt"
	"sync"

	pool "github.com/sysulq/goroutine-pool"
)

func main() {
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
