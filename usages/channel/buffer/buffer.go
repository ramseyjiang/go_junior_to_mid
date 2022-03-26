package main

import (
	"fmt"
	"runtime"
)

// Adding a buffer will provide storage capacity to the channel, and it need not be disposed of immediately.
// Since it’s buffered, the main channel will move further when we’re adding more data to the channel.
// Once the buffer is full, we need to dispose the data or else the deadlock will be achieved again.

// This func is used to prove a channel which is created has a buffer.
// It won't block, before the buffer is full.

// Using buffered channels and for range, we can read from closed channels. Since for closed channels,
// data lives in the buffer, we can still extract that data.
// When the buffer size is non-zero n, goroutine is not blocked until after buffer is full.
func channelCapacity() {
	fmt.Println("channelCapacity start")

	// Creating a channel with a buffer. A channel buffer size is 0 also called as unbuffered channel.
	c := make(chan string, 2)
	go sender(c)

	fmt.Printf("Length of channel c is %v and capacity of channel c is %v\n", len(c), cap(c))

	// In the sender(c), it has closed the channel, but it uses buffer, so below codes can also read data from buffer.
	for val := range c {
		// After buffer a channel, output is from the first buffer to the last one.
		fmt.Printf("Length of channel c after value '%v' read is %v\n", val, len(c))
	}

	fmt.Println("channelCapacity end")
}

func sender(dataChannel chan string) {
	dataChannel <- "Some Sample Data"       // len 2, cap 2
	dataChannel <- "Some Other Sample Data" // len 1, cap 2
	dataChannel <- "Buffered Channel"       // len 0, cap 2, block here.

	// If here does not have the close(c), it will have "fatal error: all goroutines are asleep - deadlock!"
	// After it add the close, it won't wait for receiving a data.
	close(dataChannel)
}

// Length of a channel is the number of values queued (unread) in channel buffer
// while the capacity of a channel is the buffer size.

// As many of you pointed out, the last value of active goroutines should be 1.

// Since send operation on a buffered channel is non-blocking (when not full), next fmt.Println statement executes.
// Go scheduler also schedule goroutines on fmt.Println statement due to blocking I/O operation,
// however, this operation is not always blocking.
// This is where the squares goroutine wake up again, runs the last iteration, prints the value in the channel
// using fmt.Println (again, this could be blocking), and dies.

/** When the buffer is full, any value sent to the channel is added to the buffer by throwing out last value
in the buffer which is available to read.
But there is a catch, read operation on buffered is thirsty.
That means, once read operation starts, it will continue until the buffer is empty.
Technically, that means goroutine that reads buffer channel will not block until the buffer is empty.
*/
func channelBufferFullBlock() {
	fmt.Println("channelBufferFullBlock start")

	// No new goroutine is triggered, so it is still one.
	fmt.Println("active goroutines", runtime.NumGoroutine()) // active goroutines 1

	c := make(chan int, 3)
	go squares(c) // trigger a new goroutine works.

	fmt.Println("active goroutines", runtime.NumGoroutine()) // active goroutines 2

	c <- 1
	c <- 2
	c <- 3
	c <- 4 // After this executes, one channel is full, so that goroutine stops working. Blocks here.
	c <- 5

	fmt.Println("How many core cpu is available?", runtime.NumCPU())

	// because one channel is full, so the full channel stops working, only one channel still works.
	fmt.Println("active goroutines", runtime.NumGoroutine()) // active goroutines 1

	go squares(c) // trigger a new goroutine works.

	fmt.Println("active goroutines", runtime.NumGoroutine()) // active goroutines 2

	c <- 6
	c <- 7
	c <- 8 // blocks here, after this line, buffer is full, it won't run the next c<-9. Because it will run c<9 at first.
	c <- 9 // why?

	fmt.Println("active goroutines", runtime.NumGoroutine())

	fmt.Println("channelBufferFullBlock end")
}

// The for loop inside squares goroutine runs 4 iterations.
// On the fourth iteration, it blocks since the buffer is empty at that point.
// Hence, Go scheduler schedules main goroutine, and we feed another value to the buffer.
func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	// main() is a special goroutine.
	fmt.Println("active goroutines", runtime.NumGoroutine()) // active goroutines 1
	// channelCapacity()
	channelBufferFullBlock()
}
