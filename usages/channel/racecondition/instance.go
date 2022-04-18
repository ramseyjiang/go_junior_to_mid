package main

import (
	"fmt"
	"sync"
)

var j = 0

// goroutine increment global variable j
func worker(wg *sync.WaitGroup) {
	j++
	wg.Done()
}

// At any condition, you shouldn’t rely on Go’s scheduling algorithm and implement your own logic to synchronize different goroutines.
// main is a special goroutine. It won't be 1000, because the program has a race condition.
// One way to make sure that only one goroutine complete all 3 above steps at a time is by implementing the mutex.
// Please go to mutex folder to see how to do it in the correct way.
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	// wait until all 1000 goroutines are done
	wg.Wait()

	// When the main goroutine starts executing after wg.Wait() call, we are printing j.
	// It won't be 1000, because the program has a race condition.
	fmt.Println("value of j after 1000 operations is ", j)
}
