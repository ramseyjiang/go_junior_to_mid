package main

import (
	"fmt"
	"strings"
)

func main() {
	replacer()
	compare()
	contains()
	index()
}

// The strings.NewReplacer is used to take arguments with a string to replace ("a"), and a string to replace it with ("o")
func replacer() {
	str := "Ga racks!"

	// define a replacer first.
	strReplacer := strings.NewReplacer("a", "o")

	// execute replace after input a string which should be replaced.
	// When we pass a string to the Replacer value’s Replace method, it returns a string with those replacements made.
	fixed := strReplacer.Replace(str)
	fmt.Println(fixed) // Go rocks!
}

// strings.Compare is used to compare two strings, and return -1, 0, 1.
func compare() {
	str1 := "abc"
	str2 := "abd"
	str3 := "abc"

	a := strings.Compare(str1, str2)
	b := strings.Compare(str1, str3)
	c := strings.Compare(str2, str1)

	fmt.Println(a, b, c)
}

// strings.Contains use to check whether a substring is in the string or not.
func contains() {
	str := "abc"
	str1 := "ab"
	str2 := "ac"

	a := strings.Contains(str, str1)
	b := strings.Contains(str, str2)
	fmt.Println(a, b)
}

func index() {
	str := "abc"
	str1 := "ab"
	str2 := "ac"
	a := strings.Index(str, str1)
	b := strings.Contains(str, str2)
	fmt.Println(a, b)
}
