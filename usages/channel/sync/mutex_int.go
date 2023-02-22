package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	plus := func() {
		lock.Lock()         // request exclusive use of the critical section—in this case the count variable— guarded by a Mutex, lock.
		defer lock.Unlock() // indicate that we’re done with the critical section lock is guarding.
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	minus := func() {
		lock.Lock()         // request exclusive use of the critical section—in this case the count variable— guarded by a Mutex, lock.
		defer lock.Unlock() // indicate that we’re done with the critical section lock is guarding.
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	// Increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			plus()
		}()
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			minus()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}
