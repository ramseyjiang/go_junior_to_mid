package main

import "fmt"

// Create channels which are unidirectional in nature.
// For example, receive-only channels which allow only read operation on them and
// send-only channels which allow only to write operation on them.

// The roc is receive-only channel as arrow direction in make function points away from chan keyword.
// While soc is send-only channel where arrow direction in make function points towards chan keyword.

// Using unidirectional channels increases the type-safety of a program. Hence, the program is less prone to error.

func main() {
	roc := make(<-chan int)
	soc := make(chan<- int)

	fmt.Printf("Data type of roc is `%T`\n", roc)
	fmt.Printf("Data type of soc is `%T`\n", soc)

	c := make(chan string)
	go greet(c)
	c <- "Davy"

	// roc <- "Ramsey" // If uncommented this line, it will notice an error directly. Because it is a read only channel.
}

// A receive-only channel example.
func greet(roc <-chan string) {
	fmt.Println("Hello " + <-roc + "!")
}
