package main

import (
	"fmt"
)

type Liters float64
type Gallons float64
type Milliliters float64

/**
Liters' method and Milliliters' method have the same name, if method name belongs to different types, it works.
That's the different between method and function in Go.
The other different is that when define a method, it includes a receiver parameter.
The method becomes associated with the type of the receiver parameter.
From then on, that method can be called on any value of that type.
*/

// To define a method, provide a receiver parameter in parentheses before the method name.
// (l Liters) is the receiver parameter.

// ToGallons method belongs to Liters Type.
func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

// ToGallons method belongs to Milliliters Type.
func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func main() {
	soda := Liters(2)
	// soda.ToGallons() will convert Liters to Gallons.
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())

	water := Milliliters(500)
	// water.ToGallons() will convert Milliliters to Gallons.
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())

	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToMilliliters())
}
