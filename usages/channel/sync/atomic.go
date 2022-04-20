package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func addOne(counter *int32, wg *sync.WaitGroup) {
	defer wg.Done()

	// AddInt32, passing a pointer to the counter and a delta as parameters.
	atomic.AddInt32(counter, 1)
}

// sync/atomic is provided by Go out of the box and allows you to safely perform atomic operations with integers in multiple goroutines.
// Please click atomic, show more details.
// It splits an algorithm into 100 smaller pieces and dispatched them in batches of two that increment the counter by one asynchronously.
// This approach is based on sharing memory — more specifically, sharing a reference to an integer — and synchronizing the access to it.
func main() {
	thisCounter := int32(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(2)

		go addOne(&thisCounter, &wg)
		go addOne(&thisCounter, &wg)
	}

	//  The goroutines are running in the background and will still be running when log.Printf is executed,
	//  so we need to wait for them by using a sync.WaitGroup.
	wg.Wait()

	// Safely access the counter value, we can use LoadInt32.
	log.Printf("Counter: %d", atomic.LoadInt32(&thisCounter))
}
