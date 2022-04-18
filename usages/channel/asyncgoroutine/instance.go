package main

import (
	"fmt"
)

func main() {
	var numbers []int // nil
	done := make(chan struct{})
	// start an anonymous goroutine to initialise array
	go func() {
		numbers = make([]int, 2)
		done <- struct{}{}
	}()

	// do something synchronous, if it does not has "<-done", it will has the below panic. Otherwise, use the waitgroup for sync.
	// panic: runtime error: index out of range [0] with length 0
	<-done                  // read done from channel
	numbers[0] = 1          // will not panic anymore
	fmt.Println(numbers[0]) // 1
}
