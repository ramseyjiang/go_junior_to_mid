package main

import (
	"fmt"
	"math"
)

/*
	The parenthensises includes types, all types in it, they are return values type.
	The function returns an integer, a boolean and a string.
	If it makes the purpose of the return values clearer, you can supply names for each one. Please check floatParts func.
*/
func manyReturns() (int, bool, string) {
	return 1, true, "hello"
}

/*
	intergerPart is the name for the first return value.
	fractionalPart is the name for the second return value.
*/
func floatParts(number float64) (intergerPart int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func main() {
	// Store each return value in a variable.
	myInt, myBool, myString := manyReturns()
	fmt.Println(myInt, myBool, myString) // 1 true hello

	cans, reminder := floatParts(1.26)
	fmt.Println(cans, reminder) // 1 0.26
}
