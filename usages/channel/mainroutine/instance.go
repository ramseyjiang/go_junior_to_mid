package main

import (
	"fmt"
	"sync"
)

// go tool compile -S instance.go >> instance.C
// The above is the command generate disassembly file
func main() {
	firstCase()
	correctCase()
}

func firstCase() {
	var wg = &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// fmt.Println("main goroutine is run at ", time.Now())
		go func() {
			fmt.Println(i)
			// fmt.Println("go goroutine is run at ", time.Now())
			wg.Done()
		}()
	}

	// Using a WaitGroup to prevent the main() Goroutine exiting before the loop finishes.
	wg.Wait()
}

func correctCase() {
	var wg = &sync.WaitGroup{}
	fmt.Println("-----correctCase start-----")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("-----correctCase end-----")
}
