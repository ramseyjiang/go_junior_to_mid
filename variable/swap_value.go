package main

import "fmt"

func main() {
	/* define  */
	var a int32 = 100
	var b int32 = 200

	fmt.Printf("Before swap, a value is : %d\n", a)
	fmt.Printf("Before swap, b value is : %d\n", b)

	/* Invoke swapValue()
	 * &a point to a, it is the a's address.
	 * &b point to b, it is the b's address.
	 */
	swapValue(&a, &b)

	fmt.Printf("After swap, a value is : %d\n", a)
	fmt.Printf("After swap, b value is : %d\n", b)
}

func swapValue(x *int32, y *int32) {
	var temp int32
	temp = *x /* temp is used to store the value which is on x's address */
	*x = *y   /* Let the value which is on y's address, to x's address*/
	*y = temp /* Give temp value to y's address */
}
