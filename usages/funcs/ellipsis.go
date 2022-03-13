package main

import (
	"fmt"
	"math"
)

/*
	"..." is called ellipsis in English.

	To declare a variadic function, place an ellipsis (...) before the type of the last parameter in the function
	declaration. That parameter will then receive all the variadic arguments as a slice.

	Only the last parameter in a function definition can be variadic.
	If a function does not have "...", when use it, it should pass the exact arguments.
*/
func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

// Here uses "...", so it can get a maximum number in unlimited numbers.
func maximum(numbers ...float64) float64 {
	// start with a very low value
	max := math.Inf(-1)

	// for ... range, is similar with PHP foreach.
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}

	return result
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func main() {
	severalInts()        // []
	severalInts(1)       // [1]
	severalInts(1, 2, 3) // [1 2 3]

	/*
		When calling a variadic function, you can use a slice in place of the variadic arguments by typing
		an ellipsis after the slice. If not, it will have an error.
	*/
	intSlice := []int{1, 2, 3} // it is the same as php define an int array
	severalInts(intSlice...)
	stringSlice := []string{"a", "b", "c", "d"} // it is the same as php define a string array.
	mix(1, true, stringSlice...)

	fmt.Println(maximum(82, 91, 65, 72))
	fmt.Println(inRange(1, 100, -12.5, 3.2, 0, 50, 103.5)) // [3.2 50]
}
