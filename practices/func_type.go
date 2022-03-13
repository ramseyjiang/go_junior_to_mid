package main

import "fmt"

func callFunction(passedFunction func()) {
	passedFunction()
}

func callTwice(passedFunction func()) {
	passedFunction()
	passedFunction()
}

func callWithArguments(passedFunction func(string, bool)) {
	passedFunction("This sentence is", false)
}

func printReturnValue(passedFunction func() string) {
	fmt.Println(passedFunction())
}

func functionA() {
	fmt.Println("functionA is called")
}

func functionB() string {
	fmt.Println("functionB is called")
	return "Returning from function"
}

func functionC(a string, b bool) {
	fmt.Println("functionC is called")
	fmt.Println(a, b)
}

/**
The result is as below:

functionA is called
functionA is called
functionA is called
functionC is called
This sentence is false
functionB called
Returning from function
*/
func main() {
	callFunction(functionA)
	callTwice(functionA)
	callWithArguments(functionC)
	printReturnValue(functionB)
}
