package main

import "fmt"

func main() {

	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("The first row is - value c is %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("The second row is - value c is %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("The third row is - value c %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("The fourth row is - value c %d\n", c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("The fifth row is - value c %d\n", c)
}
