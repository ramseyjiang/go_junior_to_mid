package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func printNumbersParallel(wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	// Parallel execution (side by side)
	go printNumbersParallel(&wg)
	go printNumbersParallel(&wg)

	wg.Wait()
	fmt.Println()
}
