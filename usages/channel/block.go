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
Hence, this read operation will be blocking. Similarly, if you are to send data to a channel,
it will block current goroutine and unblock others until some goroutine reads the data from it. Hence,
this sends operation will be blocking.
*/
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go abc(channel1)
	fmt.Println("---channel block line---") // It will run fmt.Print("channel block line") before go channel run.
	go def(channel2)

	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println() // If not has the last fmt.Println(), the result will be "adbecf%". After has it, result is "adbecf"
}
