package main

import (
	"fmt"
	"sync"
)

/**
WaitGroup is opposite to select where you needed only one condition to be true,
but here you need all conditions to be true in order to unblock the main goroutine.

WaitGroup is a struct with a counter value which tracks how many goroutines were spawned and how many have completed their job.
This counter when reaches zero, means all goroutines have done their job.
*/

func main() {
	service := func(wg *sync.WaitGroup, i int) {
		fmt.Printf("ID:%d: Hello goroutines!\n", i)

		// Here we call Done using defer keyword to ensure that before we exit the goroutine’s closure,
		// we indicate to the WaitGroup that we’ve exited.
		defer wg.Done()
	}

	var wg sync.WaitGroup // create WaitGroup empty struct

	goRoutines := 5
	wg.Add(goRoutines)

	// After for loop has done executing, it still did not pass control to other goroutines.
	// This is done by calling Wait method on wg like wg.Wait().
	// This will block the main goroutine until the counter reaches 0.
	// Once the counter reaches 0 from goroutines, main goroutine will unblock and starts executing further code.
	for i := 0; i < goRoutines; i++ {
		// Add method accept type of int, that means delta can also be negative.
		// That will be used when you know amount of goroutines, and then you can try to use -1 to make it work.
		go service(&wg, i+1)
	}

	// Here we call Wait, which will block the main goroutine until all goroutines have indicated they have exited.
	wg.Wait()

	fmt.Println("main goroutine finished, all goroutines have finished")
}
