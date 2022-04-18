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

// Using buffered channels and for range, we can read from closed channels.
// Since for closed channels, data lives in the buffer, we can still extract that data.
func channelCapacity() {
	fmt.Println("channelCapacity start")

	// Creating a channel with a buffer. A channel buffer size is 0 also called as unbuffered channel.
	c := make(chan string, 2)
	go sender(c)

	fmt.Printf("Length of channel c is %v and capacity of channel c is %v\n", len(c), cap(c))

	// In the sender(c), it has closed the channel, but it uses buffer, so below codes can also read data from buffer.
	for val := range c {
		// After buffer a channel, output is from the first buffer to the last one.
		fmt.Printf("Length of channel c has read value '%v' length is %v, capacity is %v\n", val, len(c), cap(c))
	}

	fmt.Println("channelCapacity end")
}

func sender(dataChannel chan string) {
	dataChannel <- "Some Sample Data"       // len 2, cap 2
	dataChannel <- "Some Other Sample Data" // len 1, cap 2
	dataChannel <- "Buffered Channel"       // len 0, cap 2, block here.

	// If here does not have the close(c), it will have "fatal error: all goroutines are asleep - deadlock!"
	// After it adds the close, it won't wait for receiving a data.
	close(dataChannel)
}

// Length of a channel is the number of values queued (unread) in channel buffer while the capacity of a channel is the buffer size.

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

	// No new goroutine is triggered, so it is still one. The main is a special goroutine.
	fmt.Println("active goroutines", runtime.NumGoroutine()) // active goroutines 1

	c := make(chan int, 3)
	go squares(c) // trigger a new goroutine works.

	// active goroutines 2, because the main one and the go squares(c) are running.
	fmt.Println("active goroutines", runtime.NumGoroutine())

	fmt.Println("How many core cpu is available?", runtime.NumCPU())

	fmt.Printf("channel c, length of is %v, capacity is %v\n", len(c), cap(c))
	c <- 0
	fmt.Printf("channel c, length of is %v, capacity is %v\n", len(c), cap(c))
	c <- 1
	fmt.Printf("channel c, length of is %v, capacity is %v\n", len(c), cap(c))
	c <- 2
	fmt.Printf("channel c, length of is %v, capacity is %v\n", len(c), cap(c))
	c <- 3                                                   // After this executes, one channel is full, so that goroutine squares stops working. squares will output results in loop.
	fmt.Println("active goroutines", runtime.NumGoroutine()) // active goroutines 1

	fmt.Println()
	fmt.Println("After func squares outputs all variables, channel c is empty again. So it can receive values again.")

	// In squares func, if it has close(c), only c <- 4 will run successfully.
	// If it does not have close(c), c <- 4, c <- 5 and c <- 6 will run successfully.
	// Because channel c has 3 buffer, it can put 3 values in buffer after it is closed.
	fmt.Printf("channel c, length of is %v, capacity is %v\n", len(c), cap(c))
	c <- 4
	fmt.Printf("channel c, length of is %v, capacity is %v\n", len(c), cap(c))

	c <- 5
	fmt.Printf("channel c, length of is %v, capacity is %v\n", len(c), cap(c))
	c <- 6
	fmt.Printf("channel c, length of is %v, capacity is %v\n", len(c), cap(c))

	// fmt.Println("active goroutines", runtime.NumGoroutine()) // active goroutines 1
	// go squares(c) // trigger another new goroutine works.
	// fmt.Println("active goroutines", runtime.NumGoroutine()) // active goroutines 2
	// c <- 7 // After this executes, one channel is full, so that goroutine squares stops working. squares will output results in loop.
	// c <- 8 // If line 102 is uncommented, c <- 8 and c <- 9 will be store in the buffer.
	// c <- 9

	// because the channel is full, so the full channel stopped working, only main channel still works.
	fmt.Println("active goroutines", runtime.NumGoroutine()) // active goroutines 1

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
	// If it has close(c), after the loop finishes, channel c will be closed. Then channel c 3 buffers will invalid.
	// If it does not have close(c), the channel c 3 buffers will always valid.
	// close(c)
}

func main() {
	// main() is a special goroutine.
	fmt.Println("active goroutines", runtime.NumGoroutine()) // active goroutines 1
	channelBufferFullBlock()
	channelCapacity()
}
