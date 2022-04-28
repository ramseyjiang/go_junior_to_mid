package main

import "fmt"

/*
Naked returns are okay if the function is a handful of lines.
Once it’s a mid-sized function, be explicit with your return values.
Corollary: it’s not worth it to name result parameters just because it enables you to use naked returns.
Clarity of docs is always more important than saving a line or two in your function.

Don’t name result parameters just to avoid declaring a var inside the function
*/
func main() {
	sum := addAndPrint(1, 2)
	fmt.Println(sum)
}

func addAndPrint(a int, b int) (sum int) {
	sum = a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)
	return
}
