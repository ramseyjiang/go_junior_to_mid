package main

import (
	"fmt"
)

// When you call a method that requires a pointer receiver on a variable with a none pointer type,
// Go will automatically convert the receiver to a pointer for you.
// The same is true for variables with pointer types; if you call a method requiring a value
// receiver, Go will automatically get the value at the pointer for you and pass that to the method.

// The type name begins with uppercase letter, it can be exported from outside.
// The type name begins with lowercase letter, it cannot be exported from outside
type nameType string
type ExampleType string
type Number int

// To define a method, provide a receiver parameter in parentheses before the method name.
// (n nameType) is the receiver parameter.
// In between the keyword func and the name of the function we've added a “receiver”.
// The receiver is like a parameter – it has a name and a type – but by creating the function in this way
// it allows us to call the function using the "." operator.

// n is the receiver parameter value, the name of the receiver parameter in the method is not important than the type.
// nameType is the receiver parameter type. It is the value to call method on.
// The value you’re calling the method on is known as the method receiver.
// sayHi has a receiver parameter with a type of nameType, you can call the sayHi() method on any nameType value.
// Once a method is defined on a type, it can be called on any value of that type.
// By convention, Go developers usually use a name consisting of a single letter—the first letter of the receiver’s
// type name, in lowercase.
// Go uses receiver parameters instead of the “self” or “this” values seen in other languages.
func (n nameType) sayHi() {
	// Print the receiver parameter's value. Go lets you name a receiver parameter whatever you want.
	fmt.Println("Hi from", n)
}

// Create two different nameType values
func compareTwoTypeValues() {
	// The receiver is value.
	// The receiver parameter is nameType.
	value := nameType("a nameType value")
	value.sayHi()

	// The receiver is anotherValue.
	// The receiver parameter is nameType.
	anotherValue := nameType("another value")
	anotherValue.sayHi()
}

// An example includes define method parameters and multiple returns.

func (e ExampleType) MethodWithParametersAndReturn(number int, flag bool) (int, bool) {
	fmt.Printf("Receiver name is \"%s\", int parameter value is %v, bool parameter value is %v\n",
		e, number, flag)
	return number, flag
}

func displayExampleTypeWithParameters() {
	value := ExampleType("ExampleType value")
	a, b := value.MethodWithParametersAndReturn(4, true)
	fmt.Println(a, b)
}

// Double is used to double a number.
// *Number means change the receiver parameter to a pointer type.
func (n *Number) Double() {
	*n *= 2 // Update the value at the pointer.
}

func doubleSomeNumber() {
	/*
		Number(3).Double() is the error way.
		Because if you try to call a method on the value itself, Go won’t be able to get a pointer.
		So separately them, firstly named a receiver parameter, and then call the Double() method. It will work.
	*/
	number := Number(3)
	fmt.Println("Original value of number:", number)
	number.Double()                                            // We don't have to update the method call, we just invoke the Double() method.
	fmt.Println("number after calling Double method:", number) // Value at pointer was updated.
}

func main() {
	compareTwoTypeValues()
	displayExampleTypeWithParameters()
	doubleSomeNumber()
}
