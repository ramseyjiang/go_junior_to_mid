package main

import "fmt"

/*
Three ways define variable in Go
Println, it will auto change a new line during it prints.
Printf, it won't auto change a new line, %v means value, %T means type.
In go, if a variable is declared, but it won't use in your code, it will have an error.
When using an acronym, it should all uppercase. example: HTTP, URL
*/
func main() {
	var i int
	i = 32
	fmt.Println(i)

	var j int32 = 33
	fmt.Println(j)

	k := 23
	/* := means create a new variable and give a value to the new value.
	// var k int = 23
	// k := 23	// If code has line 17, and also has line 18, it will have an error.
	Because k have been defined by line 7, the line 18 cannot redefine it again. If move the ":" in line 18, it will no error anymore.
	*/
	fmt.Println(k)

	fmt.Printf("%v, %T", k, k)
	fmt.Println(" \n ")

	m := 24.81
	fmt.Printf("%v, %T", m, m)
	fmt.Println(" \n ")

	var n float32 = 23.63
	fmt.Printf("%v, %T", n, n)

	HTTP := "https://google.com"
	fmt.Println(HTTP)

	var r = 'a' // rune is a special data type in go, it almost the same with int32
	fmt.Printf("%v, %T\n", r, r)
}
