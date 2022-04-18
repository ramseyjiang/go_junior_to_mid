package main

import "fmt"

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

/**
go run base/usages/channel/block.go, result is: "adbecf".

It is not abcdef, because the channel block is used in this case.

Channels ensure the sending goroutine has sent the value before the receiving channel attempts to use it.
Channels do this by blocking—by pausing all further operations in the current goroutine.
A send operation blocks the sending goroutine until a goroutine executes a received operation on the same channel.

The abc goroutine blocks each time it sends a value to a channel until the main goroutine receives from it. The
def goroutine does the same. The main goroutine becomes the orchestrator of the abc and def goroutines, allowing
them to proceed only when it’s ready to read the values they’re sending.
*/

/**
If you are trying to read data from a channel but channel does not have a value available with it,
it blocks the current goroutine and unblocks other in a hope that some goroutine will push a value to the channel.
Hence, this read operation will be blocked.
Similarly, if you send data to a channel, it will block current goroutine and unblock others until some goroutine reads the data from it.
Hence, this sends operation will be blocked.
*/

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go abc(c1)
	fmt.Println("---channel block line---") // It will run fmt.Print("channel block line") before go channel run.
	go def(c2)

	fmt.Print(<-c1)
	fmt.Print(<-c2)
	fmt.Print(<-c1)
	fmt.Print(<-c2)
	fmt.Print(<-c1)
	fmt.Print(<-c2)
	fmt.Println() // If not has the last fmt.Println(), the result will be "adbecf%". After has it, result is "adbecf"

	// The below code will be triggered deadlock. Because all read data before, they are reading data from the channel
	// which is scheduled by the Go Scheduler.
	// All read operations <-c1 or <-c2 are non-blocking because data is present in channel c1 and c2 to be read from.
	// However, the below one, it does not have any data to be read from, so it triggers the deadlock.
	// fmt.Println(<-c1)
}
