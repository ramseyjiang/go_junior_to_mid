package main

import (
	"fmt"
	"sync"
)

func main() {
	readClosedChan()
	closedChanUnblockAllGoroutines()
}

func readClosedChan() {
	intStream := make(chan int)
	close(intStream)

	// Notice: The channel is never placed anything before, the int channel default value is 0.
	// ok value is false, it indicates the channel was never placed any value in it.
	// The value from a closed channel, is always there.
	integer1, ok1 := <-intStream
	integer2, ok2 := <-intStream
	fmt.Printf("(%v): %v", ok1, integer1)
	fmt.Printf("(%v): %v", ok2, integer2)
}

// closedChanUnblockAllGoroutines is an example of closed a channel, it will unblock all blocked goroutines about this channel.
func closedChanUnblockAllGoroutines() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// The begin channel is blocked at here, because it does not receive any value before.
			// the goroutine waits until it is told it can continue.
			<-begin

			// because the begin channel is blocked, so until the begin channel is unblocked, the fmt will be executed.
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("\nUnblocking goroutines...")

	// close the channel, thus unblocking all the goroutines simultaneously.
	close(begin)

	wg.Wait()
}
