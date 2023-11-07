package main

import (
	"fmt"
	"sync"
)

var (
	sharedInt = 0
	mutex     sync.Mutex
)

func worker(wg *sync.WaitGroup) {
	mutex.Lock() // Lock the mutex before updating sharedInt
	sharedInt++
	mutex.Unlock() // Unlock the mutex after updating
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()

	fmt.Println("value of sharedInt after 1000 operations is ", sharedInt)
}
