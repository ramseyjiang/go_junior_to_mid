package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbersConcurrently(wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Concurrent execution (interleaved)
	go printNumbersConcurrently(&wg)
	go printNumbersConcurrently(&wg)

	wg.Wait()
	fmt.Println()
}
