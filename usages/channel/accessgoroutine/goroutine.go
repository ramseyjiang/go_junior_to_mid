package main

import "fmt"

// Goroutines allow for concurrency: pausing one task to work on others. And in some situations they allow parallelism.

// Channel two features are below.
// 1.Channels can be sent as parameters to different goroutines
// 2.Channels work as both a publisher and subscriber model.
// There is a one-to-one relationship between the sender and the receiver of the data.
// If there’s no receiver, a message is stuck with the sender.

// Go statements can’t be used with return values, because there’s no guarantee the return value will be ready before we attempt to use it.

// The only practical way to use a channel is to communicate from one goroutine to another goroutine.
// Not only does channel allow you to send values from one goroutine to another, they ensure the sending goroutine
// has sent the value before the receiving goroutine attempts to use it.

// Channels do this by blocking—by pausing all further operations in the current goroutine.
// A send operation blocks the sending goroutine until a goroutine executes a received operation on the same channel.

// To send a value on a channel, use the <- operator.  "myChannel <- 3.14", send 3.14 to myChannel.
// use the <- operator to receive values from a channel.  "<- myChannel", receive values from myChannel.

// Go programs end when the main goroutine stops, even if other goroutines have not completed their work yet.

// A channel can be closed so that no more data can be sent through it.

// The greeting will include third, fifth and seventh outputs.
func greeting(exampleChannel chan string, content string) { // Take a channel as a parameter
	exampleChannel <- content // Send a value to the channel

	fmt.Println(exampleChannel, content)
}

func firstWayCreateChannel() chan string {
	firstWayChannel := make(chan string) // Create a new channel and declare a variable at once.

	return firstWayChannel
}

func secondWayCreateChannel() chan string {
	// It can be used to declare a variable to hold a channel, but it is only a nil channel, cannot be used as a real channel.
	// var secondWayChannel chan string

	// created an object of type channel that can be used to transfer string data within the goroutines.
	secondWayChannel := make(chan string)

	return secondWayChannel
}

// TriggerGoroutine Check the output sort, you will know during the channel read and write, how the channel block works.
func main() {
	firstWayChannel := firstWayCreateChannel()

	// This output is the fifth. Pass the channel as a parameter in a new goroutine
	go greeting(firstWayChannel, "firstWayChannel")

	// Telling the main function to keep waiting until the channel receives the data.
	fmt.Println(firstWayChannel, "first print")

	/**
	When write or read data from a channel, that goroutine is blocked and control is passed to available goroutines.
	What if there are no other goroutines available, imagine all of them are sleeping.
	That’s where deadlock error occurs crashing the whole program.

	Receiver goroutine can find out the state of the channel using val,
	ok := <- channel syntax where ok is true if the channel is open or read operations can be performed
	and false if the channel is closed and no more read operations can be performed.
	A channel can be closed using close built-in function with syntax close(channel).
	*/

	secondWayChannel := secondWayCreateChannel()

	// Store the received value in a receiveFirstValue, and store the status in ok.
	receiveFirstValue, ok := <-firstWayChannel

	fmt.Println(receiveFirstValue, ok) // second print

	go greeting(secondWayChannel, "secondWayChannel")
	receiveSecValue, ok := <-secondWayChannel
	fmt.Println(receiveSecValue, ok, "fifth print") // Receive a value from the channel

	// Channels by default are pointers. It is the sixth output.
	thirdChannel := make(chan string)
	fmt.Printf("type of third channel is %T, value of it is %v\n", thirdChannel, thirdChannel)

	go greeting(thirdChannel, "test close channel")
	fmt.Println(<-thirdChannel) // It is the eighth output

	// Close the channel. If not close, only after write into/read from the channel, the channel will asleep.
	// Closing the channel does not block the current goroutine unlike reading or writing a value to the channel.
	close(thirdChannel)

	/**
	If the below code run without close on the above, it will have the below error.
	After the value write into the channel, the channel will be asleep.
	So below codes will have an error. 	"fatal error: all goroutines are asleep - deadlock!"

	If the below code run with close on the above, it will have the error as following:
	goroutine 1 [running]:
	main.main()
	        /Users/daweijiang/go/src/golang_learn/base/usages/channel/goroutine.go:86 +0x486
	exit status 2
	*/
	// thirdChannel <- "close"	// thirdChannel has been closed on line 93, so it cannot receive a value.
	fmt.Println(thirdChannel) // Print the thirdChannel address.
}
