package main

import (
	"fmt"
	"reflect"
)

// cannot use shorthand notation with constant define. "const a := 1" is always wrong.
func main() {
	var quantity = 32
	var length, width = 1.2, 2.4
	bookName := "head first go"
	trueString := "True"
	area := length * width

	/*
		If the name of a variable, function, or type begins with a capital
		letter, it is considered exported and can be accessed from greeting
		outside the current one. (This is why the P in fmt.Println is
		capitalized: so it can be used from the main package or any other.)
	*/
	fmt.Println(reflect.TypeOf(quantity))                     // int
	fmt.Println("Length is ", length, reflect.TypeOf(length)) // float64
	fmt.Println("Width is ", width, reflect.TypeOf(width))    // float64
	fmt.Println("Area is ", area, reflect.TypeOf(area))       // float64
	fmt.Println(reflect.TypeOf(true))                         // bool
	fmt.Println(reflect.TypeOf(trueString))                   // string
	fmt.Println(reflect.TypeOf(bookName))                     // string

	myInt := 2
	myFloat := 3.1
	// Type convert
	fmt.Println(reflect.TypeOf(myInt))          // int
	fmt.Println(reflect.TypeOf(float64(myInt))) // int convert to float64
	fmt.Println(reflect.TypeOf(myFloat))        // float64
	fmt.Println(reflect.TypeOf(myInt))          // float64 convert to int
	// int and float convert to string, go to type_convert.go
}
