package main

import "fmt"

/*
and, or, xor, not, >>, <<, are called Bitwise operations
*/
func main() {
	a := 8 // 二进制 1000
	b := 3
	var c int8 = 4

	fmt.Println(a / b)      // Int divide int, the result is int
	fmt.Println(int(c) / b) // Only the same type number can do calculate
	fmt.Println(a << 3)     // equal to 64, it means 左移3位 1 000 000
	fmt.Println(a << 5)     // equal to 128, it means 左移4位 100 000 000 = 256
	fmt.Println(a >> 3)     // equal to 1, it means 右移3位 1000 -> 1,
	fmt.Println(a >> 4)     // equal to 1, it means 右移4位 1000 -> 0

	m := 1.5
	n := 0.3
	fmt.Printf("%v, %T\n", m/n, m/n) // The output is 5, but its type is float64.
}
