package main

import "fmt"

/**
when we write or read data from a channel, that goroutine is blocked and control is passed to available goroutines.
What if there are no other goroutines available, imagine all of them are sleeping.
That’s where deadlock error occurs crashing the whole program.

If you are trying to read data from a channel but channel does not have a value available with it,
it blocks the current goroutine and unblocks other in a hope that some goroutine will push a value to the channel.
Hence, this read operation will be blocking. Similarly, if you are to send data to a channel,
it will block current goroutine and unblock others until some goroutine reads the data from it.
Hence, this sends operation will be blocking.
*/

func main() {
	// Use the make keyword to create a new object. Specify what type of data is returned from the channel using string.
	c := make(chan string)

	// Telling the main function to keep waiting until the channel receives the data.
	fmt.Println(<-c)
}

/**
In the program above, we don’t have any other goroutines sending data to the channel.
Since there is no other channel available, the Golang program will encounter a deadlock waiting to receive the data.
These errors seem like all goroutines are asleep or simply no other goroutines are available to schedule.

Because it only sends data to a goroutine, but it does not have a goroutine to receive it before.
*/

/**
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /Users/daweijiang/go/src/golang_learn/base/usages/channel/instance1.go
exit status 2
*/

// Using the following code it will fix.
// func greet(c chan string) {
// 	<-c
// }
//
// func main() {
// 	c := make(chan string)
// 	go greet(c)
// 	c <- "John"
// 	close(c) // closing channel
// }
