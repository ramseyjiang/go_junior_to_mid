package main

import (
	"fmt"
)

type Appliance interface {
	TurnOn(name string)
	TurnOff()
}

type Fan string

func (f Fan) TurnOn(name string) {
	fmt.Println(name)
}

func (f Fan) TurnOff() {
	fmt.Println("Fan has been turned off.")
}

type CoffeePot string

func (c CoffeePot) TurnOn(name string) {
	fmt.Println(name)
}

func (c CoffeePot) TurnOff() {
	fmt.Println("CoffeePot has been turned off.")
}

func (c CoffeePot) Brew() {
	fmt.Println("Heating Up")
}

func main() {
	var fan Appliance = Fan("Windco Breeze")
	fan.TurnOn("Today is too hot.")
	fan.TurnOff()

	var coffeePot Appliance = CoffeePot("LuxBrew")
	coffeePot.TurnOn("Black Coffee.")
	coffeePot.TurnOff()
	// coffeePot.Brew() // It will has an error, because Brew() didn't define in the interface Appliance.
}
