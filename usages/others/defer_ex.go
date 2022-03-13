package main

import (
	"fmt"
	"log"
)

/**
You can place defer keyword before any ordinary function or method call, and Go will defer (that is, delay)
making the function call until after the current function exits.

The "defer" keyword ensures a function call takes place even if the calling function exits early by using the
return keyword.

The "defer" keyword can only be used with a function or method call, it cannot be used in for loops or assignments.
*/

func Socialize() error {
	defer fmt.Println("Test defer")

	fmt.Println("Hello")
	return fmt.Errorf(" test return error after defer, what will happen.") // return an error

	// The below codes won't run.
	// fmt.Println("The middle not the last output, after the defer.")
	// return nil
}

func greet(message string) {
	fmt.Println("Greeting: " + message)
}

/**
Call one
Call two
Call three
Greeting: Greet three
Greeting: Greet two
Greeting: Greet one
*/
func deferSort() {
	fmt.Println("Call one")
	defer greet("Greet one")

	fmt.Println("Call two")
	defer greet("Greet two")

	fmt.Println("Call three")
	defer greet("Greet three")
}

/**
The output is as below
Hello
The middle not the last output, after defer.
Test defer.	//The first output is deferred until after Socialize() exits.
*/
func main() {
	deferSort()

	err := Socialize()
	if err != nil {
		log.Fatal(err)
	}
}
