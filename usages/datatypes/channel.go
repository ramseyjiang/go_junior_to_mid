package main

import (
	"fmt"
	"runtime"
)

// gets a channel and prints the greeting by reading from a channel
func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

// gets a channel and write a channel to it
func greeter(cc chan chan string) {
	c := make(chan string)
	cc <- c
}

func main() {
	fmt.Println("active goroutines:", runtime.NumGoroutine()) // active goroutines: 1

	// make a channel `cc` of data type channel of string data type
	cc := make(chan chan string)

	// start greeter goroutine using cc channel
	go greeter(cc)

	fmt.Println("active goroutines:", runtime.NumGoroutine()) // active goroutines: 2

	// receive a channel c from greeter goroutine
	c := <-cc

	fmt.Println("active goroutines:", runtime.NumGoroutine()) // active goroutines: 1

	// start greet goroutine using c channel
	go greet(c)

	fmt.Println("active goroutines:", runtime.NumGoroutine()) // active goroutines: 2

	// send data to c channel
	c <- "Davy"

	fmt.Println("active goroutines:", runtime.NumGoroutine()) // active goroutines: 1
}
