package main

import "fmt"

/*

All scopes of int are under below.
int8	-128 ~ 127
int16 	-32768 ~ 32767
int32	-2 147 483 648 ~ 2 147 483 647
int64	-9 223 372 036 854 775 808 ~ 9 223 372 036 854 775 807


uint8 	0 ~ 255
uint16	0 ~ 65535
uint32  0 ~ 4294967295
don't have uint64
*/
func main() {
	var m uint32 = 42
	fmt.Printf("%v, %T\n", m, m)

	// var n int8 = 255 //If define a variable int8 equal to 255, it will overflow.
	// fmt.Printf("%v, %T\n", n, n)
}
