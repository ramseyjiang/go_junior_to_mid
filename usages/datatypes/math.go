package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// The math.Floor function takes a floatingpoint number, rounds it down to the nearest whole number, and returns that whole number.
	fmt.Println(math.Floor(2.75)) // 2

	// The strings.Title function takes a string, capitalizes the first letter of each word it contains
	fmt.Println(strings.Title("head first go")) // Head First Go
}
