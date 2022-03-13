package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

// The Go language supports first-class functions; that is, functions in Go are treated as “first-class citizens.”
// first-class functions, it means functions can be assigned to variables, and then called from those variables.
// A variable that holds a function needs to specify what parameters and return values that function should have.
// That variable can only hold functions whose number and types of parameters and return values match the specified type.
func main() {
	// This variable will hold a function with no parameters and no return value.
	var greeterFunction func()

	// This variable will hold a function with two int parameters and a float64 return value.
	var mathFunction func(int, int) float64

	greeterFunction = sayHi // Assign the sayHi function to the greeterFunction variable.
	mathFunction = divide   // Assign the divide function to the mathFunction variable.
	// If swap the sayHi and divide, it will have errors. Because types of them don't match.

	// Call the function stored in the variable.
	greeterFunction()               // Here, it will call sayHi() function.
	fmt.Println(mathFunction(5, 2)) // Here, it will call divide() function.
}
