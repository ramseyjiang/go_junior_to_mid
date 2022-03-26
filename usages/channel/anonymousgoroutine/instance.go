package main

import "fmt"

func main() {
	fmt.Println("Print by anonymous goroutine.")
	c := make(chan string)

	// launch anonymous goroutine
	go func(c chan string) {
		fmt.Println("Hello " + <-c + "!")
	}(c)

	c <- "Davy"
}
