package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func computeIntensively(wg *sync.WaitGroup) {
	defer wg.Done()

	result := 0
	for i := 0; i < 100000000; i++ {
		result += i
	}
	fmt.Printf("Result: %d\n", result)
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	start := time.Now()

	go computeIntensively(&wg)
	go computeIntensively(&wg)

	wg.Wait()

	fmt.Printf("Elapsed time: %v\n", time.Since(start))
}
