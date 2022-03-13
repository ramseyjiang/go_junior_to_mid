package main

import (
	"fmt"
	"golang_learn/customizepkgs/keyboard"
	"log"
)

func main() {
	fmt.Println("Enter a temperature in Fahrenheit:")

	// Call greeting() to get a temperature.
	fahrenheit, err := keyboard.GetFloat()

	if err != nil {
		log.Fatal(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}
