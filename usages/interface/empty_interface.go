package main

import "fmt"

/*
	So if declare an interface type that did not require any methods, it will be satisfied by any type.
	type Anything interface {}	//Can be exported anywhere.
	type anything interface {}	//Cannot be exported.
	The above interfaces are the empty interface.
	The empty interface does not require any methods to satisfy it, so it is satisfied by all types.
*/

// Do not try to call any methods on an empty interface value. If you want to call, it will have an error.
// You can try to uncomment all codes in the main and AcceptAnything method, then run it, you will see errors.

// Anything is an empty interface type.
type Anything interface {
}

func AcceptAnything(thing Anything) {
	fmt.Println(thing)

	// Using an empty interface, it will not be allowed to do anything with the variable.
	// fmt.Println(thing + 1)	//If uncomment, it will show an error.
}

/**
The first assert the type using myVar.(int) .
Into v is assigned the typed variable if the type assertion is successful.
And into ok is assigned a boolean value of whether the assertion was successful.


When you use assert, it means you knew ahead of time the type, make sure the type assertion to work.
Even if you switch and check for multiple options,
you still have to know of the concrete type you assert in compile time.
*/
func typeAssertion() {
	var testVar Anything = 10

	v, ok := testVar.(int) // testVar.(int) is a type assertion.
	if ok {
		fmt.Println(v)
	}
}

func main() {
	AcceptAnything(3)
	typeAssertion()
}
