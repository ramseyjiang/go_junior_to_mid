package main

import "fmt"

func sayHello() {
	fmt.Println("Hi")
}

func sayBye() {
	fmt.Println("Bye")
}

// This twiceFunc accepts another function as a parameter.
func twiceFunc(theFunction func()) { // It uses "func()" as a type declare for "theFunction".
	theFunction() // Call the passed-in function.
	theFunction() // Call the passed-in function again.
}

// In main, it calls twiceFunc and pass the sayHello function as an argument, which causes sayHello to be run twice.
// Then we call twice with the sayBye function, which causes sayBye to be run twice.
func main() {
	twiceFunc(sayHello) // Pass the "sayHello" function to the twiceFunc function.
	twiceFunc(sayBye)   // Pass the "sayBye" function to the twiceFunc function.
}
