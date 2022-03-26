package main

import (
	"fmt"
	"sync"
)

var j int // j default value is 0

// goroutine increment global variable j
func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // acquire lock
	j++
	m.Unlock() // release lock
	wg.Done()
}

// Create one mutex m and pass a pointer to it to all spawn goroutines.
func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &m)
	}

	// wait until all 1000 goroutines are done
	wg.Wait()

	// value of j, here is 1000.
	// It is different the example in the racecondition folder.
	// In that one, it has race condition. Using mutex to help us to resolve race conditions.
	// But only 1 goroutine can get to read or write value in mutex.
	// When you use Mutex, the first rule is to avoid shared resources between goroutines.
	fmt.Println("value of i after 1000 operations is ", j)
}
