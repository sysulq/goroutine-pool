package pool

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	wp := WorkerPool{
		MaxWorkersCount:       10,
		MaxIdleWorkerDuration: 60 * time.Second,
	}

	wp.Start()

	for i := 0; i < 10; i++ {
		assert.True(t, wp.Serve(func() {
			fmt.Println("hello")
		}))
	}
}

func TestStop(t *testing.T) {
	wp := WorkerPool{
		MaxWorkersCount: 10,
	}

	wp.Start()

	for i := 0; i < 10; i++ {
		assert.True(t, wp.Serve(func() {
			fmt.Println("hello")
		}))
	}

	wp.Stop()
}

func BenchmarkPool(b *testing.B) {
	num := 10000
	wp := WorkerPool{
		MaxWorkersCount: 50000,
	}
	wp.Start()
	wg := sync.WaitGroup{}

	for i := 0; i < num; i++ {
		wg.Add(1)
		wp.Serve(func() {
			wg.Done()
		})
	}
	wg.Wait()
	b.ResetTimer()

	for k := 0; k < b.N; k++ {
		wg.Add(1)
		// assert.True(b, wp.Serve(func() {
		wp.Serve(func() {
			rand.Int()
			time.Sleep(10 * time.Microsecond)
			wg.Done()
		})
		wg.Wait()
	}

}

func BenchmarkGoroutine(b *testing.B) {
	b.ResetTimer()

	wg := sync.WaitGroup{}
	for k := 0; k < b.N; k++ {
		wg.Add(1)
		go func() {
			rand.Int()
			time.Sleep(10 * time.Microsecond)
			wg.Done()
		}()
		wg.Wait()
	}
}

func BenchmarkPoolWithoutWait(b *testing.B) {
	num := 10000
	wp := WorkerPool{
		MaxWorkersCount: 50000,
	}
	wp.Start()
	wg := sync.WaitGroup{}

	for i := 0; i < num; i++ {
		wg.Add(1)
		wp.Serve(func() {
			wg.Done()
		})
	}
	wg.Wait()
	b.ResetTimer()

	for k := 0; k < b.N; k++ {
		// wg.Add(1)
		// assert.True(b, wp.Serve(func() {
		wp.Serve(func() {
			rand.Int()
			time.Sleep(1 * time.Microsecond)
			// wg.Done()
		})
		// wg.Wait()
	}

}

func BenchmarkGoroutineWithoutWait(b *testing.B) {
	b.ResetTimer()

	// wg := sync.WaitGroup{}
	for k := 0; k < b.N; k++ {
		// wg.Add(1)
		go func() {
			rand.Int()
			time.Sleep(1 * time.Microsecond)
			// wg.Done()
		}()
		// wg.Wait()
	}
}
