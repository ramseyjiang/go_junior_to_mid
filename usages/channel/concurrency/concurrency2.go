package main

import (
	"fmt"
	"sync"
	"time"
)

func computeConcurrently(wg *sync.WaitGroup) {
	defer wg.Done()

	result := 0
	for i := 0; i < 100000000; i++ {
		result += i
		time.Sleep(1 * time.Nanosecond)
	}
	fmt.Printf("Result: %d\n", result)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	start := time.Now()

	go computeConcurrently(&wg)
	go computeConcurrently(&wg)

	wg.Wait()

	fmt.Printf("Elapsed time: %v\n", time.Since(start))
}
