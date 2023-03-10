package main

import (
	"fmt"
	"runtime"
	"time"
)

// select statements can help safely bring channels together with concepts like cancellations, timeouts, waiting, and default values.

// Created the three channels that we'll need in this exercise.
// Then, we launched our receiver function in a different Goroutine.
// This Goroutine is handled by Go's scheduler and our program continues.
// We launched a new Goroutine to send the message hello to the helloCh arguments.
// Again, this is going to occur eventually when the Go's scheduler decides.
// Our program continues again and waits for a second.
// In this break, Go's scheduler will have time to execute the receiver and the first message (if it hasn't done so yet),
// so the hello! message will appear on the console during the break.
// A new message is sent over the goodbye channel with the goodbye! text in a new Goroutine,
// and our program continues again to a line where we wait for an incoming message in the quitCh argument.
// We have launched three Goroutines already–the receiver that it is still running,
// the first message that had finished when the message was handled by the select statement,
// and the second message has been printed almost immediately and had finished too.
// So just the receiver is running at this moment, and if it doesn't receive any other message in the following two seconds,
// it will handle the incoming message from the time structure.
// After channel type, print a message to say that it is quitting, send a true to the quitCh, and break the infinite loop where it was looping.
func main() {
	helloCh := make(chan string, 1)
	goodbyeCh := make(chan string, 1)
	quitCh := make(chan bool)
	go receiver(helloCh, goodbyeCh, quitCh)

	// send a string over a channel by simply calling the sendString method.
	go sendString(helloCh, "hello!")
	time.Sleep(time.Second)
	go sendString(goodbyeCh, "goodbye!")
	<-quitCh

	nilChan()
	loopSleep()
}

// The receiver take messages from both channels–the one that sends hello messages and the one that sends goodbye messages.
// It takes three channels–two receiving channels and one to send something through it.
// Then, it starts an infinite loop with the for keyword.
// This way we can keep listening to both channels forever.
func receiver(helloCh <-chan string, goodbyeCh <-chan string, quitCh chan<- bool) {
	for {
		// In the select structure, we ask the program to choose between one or more channels to receive their data.
		// Save this data in a variable and make something with it before finishing the select.
		// The select structure is just executed once; it doesn't matter if it is listening to more channels,
		// it will be executed only once and the code will continue executing.
		// If we want it to handle the same channels more than once, we have to put it in a for loop.
		select {
		// The first case takes the incoming data from the helloCh argument and saves it in a variable called msg.
		case msg := <-helloCh:
			println(msg)
		// The second case takes the incoming data from the goodbyeCh argument and saves it in a variable called msg
		case msg := <-goodbyeCh:
			println(msg)
		// It calls the time function.
		// After that, if we check its signature, it accepts a time and duration value and returns a receiving channel.
		// This receiving channel will receive a time, the value of time after the specified duration has passed.
		// Uses the channel it returns as a timeout.
		case <-time.After(time.Second * 2):
			println("Nothing received in 2 seconds. Exiting")
			quitCh <- true
			break
		}
	}
}

func sendString(ch chan<- string, s string) {
	ch <- s
}

// nilChan is an instance with nil chan, it is used to indicate a channel never become ready.
func nilChan() {
	var c <-chan int
	select {
	// This case statement will never become unblocked because we’re reading from a nil channel.
	case <-c:
	// when all the channels are blocked, but you also can’t block forever, you may want to time out.
	// Go’s time package provides an elegant way to do this with channels that fits nicely within the paradigm of select statements.
	// The time.After function takes in a time.Duration argument and returns a channel
	// that will send the current time after the duration you provide it.
	// This offers a concise way to time out in select statements.
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out.")
	}
}

func loopSleep() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()
	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		// Simulate work, the above select won't block.
		// select {} means select statements with no case clauses. It will block forever.
		workCounter++
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
}
