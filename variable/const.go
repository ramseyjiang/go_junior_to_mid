package main

import (
	"fmt"
)

// When define a const, the first char must be uppercase.
func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	a, b, c := 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("area is : %d", area)
	println()
	println(a, b, c)
}