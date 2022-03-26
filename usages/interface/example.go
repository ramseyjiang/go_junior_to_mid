package main

import (
	"fmt"
	"strings"
)

/**
An interface is a set of methods certain values are expected to have.
Any type that has all the methods listed in an interface definition is said to satisfy that
interface.
A type that satisfies an interface can be assigned to any variable or function parameter
that uses that interface as its type.

A concrete type specifies not only what its values can do (what methods you can call on them), but also what
they are: they specify the underlying type that holds the value’s data.

To satisfy an interface, a type must have all the methods the interface specifies. Method names, parameter types
(or lack thereof), and return value types (or lack thereof) all need to match those defined in the interface.

If you’ve assigned a value of a concrete type to a variable with an interface type, you can use a type assertion
to get the concrete type value back. Only then can you call methods that are defined on the concrete type
(but not the interface).

When you might use interface.
1. If there are certain methods you want to access without rewriting a lot of code, then you might need an interface.
2. If you need to mock up databases or external APIs in unit tests, then you might need an interface.
3. If you need to quickly plug different packages in and out of the code (like a logger) without a complete logic
	rewrite, then you might need an interface.
*/

type Whistle string

func (w Whistle) MakeSound(params string) {
	fmt.Println(params)
}

type Horn string

func (h Horn) MakeSound(params string) {
	fmt.Println(params)
}

// Only methods which are defined in the interface can be called by an interface type.

// NoiseMaker interface is satisfied by any type that has a MakeSound method.
type NoiseMaker interface {
	MakeSound(params string)
}

func play(n NoiseMaker, params string) {
	n.MakeSound(params)
}

// An interface has two types. They are the static type and the dynamic type.
// The static type of interface is the interface itself.

// An interface does not have a static value, rather it points to a dynamic value.
// A variable of an interface type can hold a value of a type that implements the interface.
// Its value of that type becomes the dynamic value of the interface and that type becomes the dynamic type of the interface.

// Must implement all the methods declared by the interface with exact signatures

// The dynamic type of interface also called a concrete type.

type NilInterface interface {
	Area() float64
	Perimeter() float64
}

// Here, "i interface{}" is an empty interface. We can pass any argument to it.
// Using "i.(type)" to control which part will go in the switch.
// In the explain method, it receives its dynamic value and dynamic type.
// This is the principle how the TypeAssert works.
func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("i stored string ", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("i stored int ", i)
	default:
		fmt.Println("i stored something else ", i)
	}
}

func main() {
	// The comment codes are the basic way.
	// var toy NoiseMaker
	// toy = Whistle("Toy Canary") // implicitly implemented.
	// toy.MakeSound()
	//
	// toy = Horn("Toy Blaster") // implicitly implemented.
	// toy.MakeSound()

	// The below way is that declare function parameters with interface types.
	play(Whistle("Toy Canary"), "Toy Canary")
	play(Horn("Toy Blaster"), "Toy Blaster")

	// we can see that zero value and type of the interface is nil. This is because, at this moment, we have
	// declared the variable s of type Shape but did not assign any value.
	var test NilInterface
	fmt.Println("value if test is", test)
	fmt.Printf("type of test is %T\n", test)

	explain("hello go interface")
	explain(24)
	explain(true)
}
