package channel

import (
	"fmt"
)

func TriggerAsyncGoroutine() {
	var numbers []int // nil
	done := make(chan struct{})
	// start a goroutine to initialise array
	go func() {
		numbers = make([]int, 2)
		done <- struct{}{}
	}()

	// do something synchronous, if it does not has "<-done", it will has the below panic.
	// panic: runtime error: index out of range [0] with length 0
	<-done                  // read done from channel
	numbers[0] = 1          // will not panic anymore
	fmt.Println(numbers[0]) // 1
}
