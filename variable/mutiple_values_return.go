package main

import "fmt"

func swapElement(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swapElement("Mahesh", "Kumar")
	fmt.Println(a, b)
}
