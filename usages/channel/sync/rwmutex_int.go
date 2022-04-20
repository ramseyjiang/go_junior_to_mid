package main

import (
	"log"
	"sync"
)

// Mutexes are allowed to synchronize the access to more complex data structures.
type counter struct {
	sharedInt int
	mux       sync.RWMutex // Locks/unlocks the data structure in reading mode.
}

func (c *counter) increment() {
	c.mux.Lock()

	// It is important to note that we need to unlock the mutexes after performing the operations.
	defer c.mux.Unlock()
	c.sharedInt++
}

func (c *counter) getValue() int {
	// Readers will be able to access it, but not writers.
	// Using this, you can achieve better performance in scenarios with a lot of readers and few writers.
	c.mux.RLock()

	// It is important to note that we need to unlock the mutexes after performing the operations.
	defer c.mux.RUnlock()
	return c.sharedInt
}

func increment(counter *counter, wg *sync.WaitGroup) {
	defer wg.Done()

	counter.increment()
}

func main() {
	c := counter{}
	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(2)

		go increment(&c, &wg)
		go increment(&c, &wg)
	}

	wg.Wait()

	log.Printf("Counter: %d", c.getValue())
}
