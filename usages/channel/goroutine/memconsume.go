package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		// GC runs a garbage collection and blocks the caller until the garbage collection is complete. It may also block the entire program.
		runtime.GC()
		var s runtime.MemStats
		// ReadMemStats populates m with memory allocator statistics. The returned memory allocator statistics are up-to-date as of thecall to ReadMemStats.
		runtime.ReadMemStats(&s)
		// Sys is the total bytes of memory obtained from the OS.
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup

	noop := func() {
		wg.Done()
		<-c
	}

	// The number of goroutines will be added is 1e4. Because the large numbers to asymptotically approach the size of a goroutine.
	const numGoroutines = 1e4
	wg.Add(numGoroutines)

	// measure the amount of memory consumed before creating our goroutines.
	before := memConsumed()
	fmt.Println("before", before)

	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()

	// measure the amount of memory consumed after creating our goroutines.
	after := memConsumed()
	fmt.Println("after", after)

	// These are just empty goroutines that don’t do anything, but it still gives us an idea of the number of goroutines we can likely create.
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
	// My laptop is Mac1 Max, which means that in theory I can spin up millions of goroutines without requiring swapping.
}
