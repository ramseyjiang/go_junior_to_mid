package main

import "fmt"

func main() {
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s) // this is a string, string

	fmt.Printf("%v, %T\n", s[2], s[2]) // 105, uint8. Because it should do the type convert first

	fmt.Printf("%v, %T\n", string(s[2]), string(s[2])) // i, string

	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b) // [116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103], []uint8
}
