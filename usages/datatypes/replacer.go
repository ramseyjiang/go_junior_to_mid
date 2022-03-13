package main

import (
	"fmt"
	"strings"
)

/*
The strings.NewReplacer function takes arguments with a string to
replace ("a"), and a string to replace it with ("o"), and returns a
strings.Replacer.
*/
func main() {
	broken := "Ga racks!"

	// define a replacer first.
	replacer := strings.NewReplacer("a", "o")

	// execute replace after input a string which should be replaced.
	// When we pass a string to the Replacer value’s Replace method, it returns a string with those replacements made.
	fixed := replacer.Replace(broken)
	fmt.Println(fixed) // Go rocks!
}
