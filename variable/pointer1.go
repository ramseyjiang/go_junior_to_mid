package main

import "fmt"

func main() {
	var a = 20
	var ip *int

	ip = &a

	fmt.Printf("variable a's address is : %x\n", &a)

	fmt.Printf("variable ip's pointer store address is: %x\n", ip)

	fmt.Printf("*ip varibale value is: %d\n", *ip)
}
