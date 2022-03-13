package main

import "fmt"

func main() {
	// Prompt the user to enter a number
	fmt.Print("Enter a number: ")

	var input float64
	a := &input

	// Scanf fills input with the number we enter
	_, err := fmt.Scanf("%f", a)
	if err != nil {
		return
	}
	output := input * 2
	fmt.Println(output)
}
