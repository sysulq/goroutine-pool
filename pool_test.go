package pool

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	Start()
	defer Stop()

	for i := 0; i < 10; i++ {
		Go(func() {
			fmt.Println("hello")
		})
	}
}

func TestStop(t *testing.T) {
	Start()
	defer Stop()

	for i := 0; i < 10; i++ {
		Go(func() {
			fmt.Println("hello")
		})
	}
}

func BenchmarkPool(b *testing.B) {
	num := 10000
	Start()
	defer Stop()
	wg := sync.WaitGroup{}

	for i := 0; i < num; i++ {
		wg.Add(1)
		Go(func() {
			wg.Done()
		})
	}
	wg.Wait()
	b.ResetTimer()

	for k := 0; k < b.N; k++ {
		wg.Add(1)
		// assert.True(b, wp.Serve(func() {
		Go(func() {
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
	Start()
	defer Stop()
	wg := sync.WaitGroup{}

	for i := 0; i < num; i++ {
		wg.Add(1)
		Go(func() {
			wg.Done()
		})
	}
	wg.Wait()
	b.ResetTimer()

	for k := 0; k < b.N; k++ {
		// wg.Add(1)
		// assert.True(b, wp.Serve(func() {
		Go(func() {
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
