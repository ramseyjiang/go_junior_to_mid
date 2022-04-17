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

type Heater struct {
	ut int // usual temperature
	nt int // night temperature
}

const ut = 26
const nt = 20

func New() Heater {
	return Heater{ut: ut, nt: nt}
}

func (h Heater) TurnOn(name string) string {
	fmt.Println("usual temperature is ", h.ut)
	fmt.Println("night temperature is ", h.nt)
	return name
}

func (h Heater) TurnOff() string {
	return "Heater has been turned off."
}
