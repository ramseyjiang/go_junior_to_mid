package main

import "fmt"

// A “fallthrough” statement transfers control to the first statement of the next case clause without validating the expression.
func main() {
	var i int8 = 1
	instance1(i)
	instance2(i)
}

func instance1(i int8) {
	fmt.Println("------ instance1 start ------")
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
		fallthrough
	case i < 20:
		fmt.Println("i is less than 20")
		fallthrough
	case i < 30:
		fmt.Println("i is less than 30")
	}
	fmt.Println("------ instance1 end ------")
}

// Good example for how fallthrough run.
func instance2(i int8) {
	fmt.Println("------ instance2 start ------")
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
		fallthrough
	case i > 20:
		fmt.Println("i is less than 20")
		fallthrough
	case i < 30:
		fmt.Println("i is less than 30")
	}
	fmt.Println("------ instance2 end ------")
}
