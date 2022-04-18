package main

import (
	"log"
	"sync"
)

type counter struct {
	value int
	mux   sync.RWMutex // Locks/unlocks the data structure in reading mode.
}

func (c *counter) increment() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.value++
}

func (c *counter) getValue() int {
	// Readers will be able to access it, but not writers.
	// Using this, you can achieve better performance in scenarios with a lot of readers and few writers.
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.value
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
