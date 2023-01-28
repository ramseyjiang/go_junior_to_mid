package main

import (
	"log"
	"sync"
)

var sharedInt int // sharedInt default value is 0

// goroutine increment global variable j
func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // acquire lock
	sharedInt++
	m.Unlock() // release lock
	wg.Done()
}

// Create one sync m and pass a pointer to it to all spawn goroutines.
func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(2)
		go worker(&wg, &m)
		go worker(&wg, &m)
	}

	// wait until all 100 goroutines are done
	wg.Wait()

	// value of j, here is 100.
	// It is different the example in the racecondition folder.
	// In that one, it has race condition. Using sync to help us to resolve race conditions.
	// But only 1 goroutine can get to read or write value in sync.
	// When you use Mutex, the first rule is to avoid shared resources between goroutines.
	log.Printf("value of i after 100 operations is %v", sharedInt)
}
