package main

import "fmt"

// In go, all varibales are defined, will have a default value.
func main() {
	var n bool
	fmt.Printf("%v, %T\n", n, n) // default bool is false

	var i32 int32
	fmt.Printf("%v, %T\n", i32, i32) // default i32 is 0

	var i64 int64
	fmt.Printf("%v, %T\n", i64, i64) // default i64 is 0

	var s string
	fmt.Printf("%v, %T\n", s, s) // default s is an empty string

	var f32 float32
	fmt.Printf("%v, %T\n", f32, f32) // default f32 is 0

	var f64 float64
	fmt.Printf("%v, %T\n", f64, f64) // default f64 is 0
}
