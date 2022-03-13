package main

import (
	"fmt"
)

/**
Any type that has all the methods listed in an interface definition is said to satisfy that interface.

Define an interface type using the interface keyword, followed by curly braces containing a list of method names,
along with any parameters or return values the methods are expected to have.

A type that satisfies an interface can be used anywhere that interface is called for.
A type can have methods in addition to those listed in the interface, but it mustn’t be missing any, or it
does not satisfy that interface.
*/
type myInterface interface { // Declare an interface type.
	MethodWithoutParameters()
	MethodWithParameters(float64)  // float64 is the type of parameter.
	MethodWithReturnValue() string // string is the type of return value.
}

type MyType int // Declare a type, We'll make it satisfy myInterface.

func (m MyType) MethodWithoutParameters() { // First required method
	fmt.Println("MethodWithoutParameters called.")
}

func (m MyType) MethodWithParameters(f float64) { // Second required method, with a float64 parameter.
	fmt.Println("MethodWithParameters called with", f)
}

func (m MyType) MethodWithReturnValue() string { // Third required method, with a string return value.
	return "Hi from MethodWithReturnValue."
}

func (m MyType) MethodNotInInterface() { // Even it has methods that aren't part of the interface, it still satisfy the interface.
	fmt.Println("MethodNotInInterface called.")
}

// A variable declared with an interface type can hold any value whose type satisfies that interface.
func main() {
	var test myInterface
	test = MyType(9)
	test.MethodWithoutParameters()
	test.MethodWithParameters(2.4)
	fmt.Println(test.MethodWithReturnValue())

	// If you want to invoke a method not in an interface, you should use interface type assert.
	// test.MethodNotInInterface()
}
