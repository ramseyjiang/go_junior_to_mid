package iface

import "fmt"

// This file can be separated two files, such as Fan and CoffeePot implements.

type Fan int

func (f Fan) TurnOn(name string) string {
	return name
}

func (f Fan) TurnOff() string {
	return "Fan has been turned off."
}

type CoffeePot string

func (c CoffeePot) TurnOn(name string) string {
	return name
}

func (c CoffeePot) TurnOff() string {
	return "CoffeePot has been turned off."
}

func (c CoffeePot) Brew() {
	fmt.Println("Heating Up")
}
