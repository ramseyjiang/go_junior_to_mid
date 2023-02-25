package main

import (
	"fmt"
	"sync"
)

func main() {
	oncePrint()
	onceFuncs()
}

func oncePrint() {
	var count int

	plus := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)

	// sync.Once is a type that utilizes some sync primitives internally to ensure that only one call to Do ever calls the function
	// passed in—even on different goroutines.
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(plus)
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}

// the output displays 1 and not 0?
// This is because sync.Once only counts the number of times Do is called, not how many times unique functions passed into Do are called.
func onceFuncs() {
	var count int
	plus := func() { count++ }
	minus := func() { count-- }

	var once sync.Once
	once.Do(plus)
	once.Do(minus)

	fmt.Printf("Count: %d\n", count)
}
