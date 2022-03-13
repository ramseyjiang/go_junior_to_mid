package main

import (
	"fmt"
	"log"
)

// A type that includes any value with a particular method. It sounds like an interface.
// You can create your own type that satisfies the error interface and stores the information you want.
// Compare with example.go, the NoiseMaker interface is satisfied by any type.
// Here error interface is only satisfied by string type.
type errorInstance interface {
	Error() string
}

type ComedyError string // Define a type with an underlying type of string
func (c ComedyError) Error() string {
	return string(c) // The Error method needs to return a string, so do a type conversion.
}

type OverheatError float64 // Define a type with an underlying type of float64
func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func checkTemperature(actual float64, safe float64) error {
	if excess := actual - safe; excess > 0 {
		return OverheatError(excess)
	}

	return nil
}

func main() {
	// Set up a variable with a type of "error"
	var err errorInstance

	// ComedyError satisfies the error interface, we can assign a ComedyError to the variable.
	err = ComedyError("What's a programmer's favorite beer? Logger!")
	fmt.Println(err)

	err = checkTemperature(121.379, 100)
	if err != nil {
		log.Fatal(err)
	}
}
