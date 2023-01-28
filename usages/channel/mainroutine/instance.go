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
		a := i // If here does not have a:=i, only using variable i in line 23, it won't output the correct num.
		go func() {
			fmt.Println(a) // if here uses the variable i, it won't be the immediate value. The variable i will become 10 during it's output.
			// fmt.Println("go goroutine is run at ", time.Now())
			wg.Done()
		}()
	}

	// Using a WaitGroup to prevent the main() Goroutine exiting before the loop finishes.
	wg.Wait()
}

// correctCase is using the correct way to pass variable i into a goroutine. It is easier than the first one.
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
