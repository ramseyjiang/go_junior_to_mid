package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, i, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(reportChannel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("***Sending value***")
	reportChannel <- "a" // Will block on this send while "main" is still asleep.
	fmt.Println("***Sending value***")
	reportChannel <- "b"
}

/**
Every second the goroutine is asleep, it will print an announcement that it’s still sleeping.

It starts with a reportNap function that causes the current goroutine to sleep for a specified number of seconds.
Every second the goroutine is asleep, it will print an announcement that it’s still sleeping.

It calls a send function that will run in a goroutine and send two values to a channel.
Before it sends anything, though, it first calls reportNap so its goroutine sleeps for 2 seconds.

When it runs, it will make both goroutines sleep for the first 2 seconds. Then the send goroutine wakes up and
sends its value. But it does not do anything further; the send operation blocks the send goroutine until the
main goroutine receives the value. That does not happen right away, because the main goroutine still needs to
sleep for 3 more seconds. When it wakes up, it receives the value from the channel. Only then is the send goroutine
unblocked. After that, it can send its second value.
*/
func main() {
	reportChannel := make(chan string)
	go send(reportChannel)
	reportNap("Receiving goroutine", 5) // In go, it will run this func before go channel run.
	fmt.Println(<-reportChannel, "first block")
	fmt.Println(<-reportChannel, "second block")
}

/**
go run practices/report_nap_goroutine.go, Result:
From 54 to 58, sending and receiving goroutine are asleep.
From 59 to 60, the sending goroutine wakes yp and sends a value.
From 61 to 62, the receiving goroutine is still sleeping.
From 63 to 64, the receiving goroutine wakes up and receives a value.
From 65 to 66, only then, the sending goroutine is unblocked, so it can send its second value.

Receiving goroutine 0 sleeping
sending goroutine 0 sleeping
sending goroutine 1 sleeping
Receiving goroutine 1 sleeping
Receiving goroutine 2 sleeping
sending goroutine wakes up!
***Sending value***
Receiving goroutine 3 sleeping
Receiving goroutine 4 sleeping
Receiving goroutine wakes up!
a first block
***Sending value***
b second block
*/
