package main

import "fmt"

/*
Structs are complex data types that hold other types.
For example, a laptop isn't just one thing. Rather, it's a combination of many parts. So we could represent a laptop as a struct.

Methods are like functions, but they are bound to a certain type or object, which are called receivers.
*/

func main() {
	usageNew()
	usageMake()
}

// Teacher Allocate memory for a custom struct type and return a pointer to it
type Teacher struct {
	Name   string
	Age    int
	Gender bool
}

// The new keyword in Golang is used to create a new instance of a variable, and it returns a pointer to the memory allocated.
func usageNew() {

	t := new(Teacher)

	fmt.Println(t)

	fmt.Println((*t).Gender, (*t).Name, (*t).Age)

	intPtr := new(int)
	fmt.Println(intPtr)
	fmt.Println(*intPtr)
	*intPtr = 6
	fmt.Println(*intPtr)
}

// The make keyword in Golang is used to create slices, maps, and channels, and it returns a value of the type that was created.
func usageMake() {
	// make is used to create a new slice of int type with a length of 2, and the slice is stored in the s variable.
	numbers := make([]int, 2)
	ages := make(map[string]int)
	fmt.Println(numbers)
	// setting the values of the slice
	numbers[0] = 1
	numbers[1] = 2

	fmt.Println(ages)
}

/*
1. new returns a pointer to the memory allocated, while make returns the value of the type being created.
2. new only works with basic types such as int, float, bool, etc. make is used for creating slices, maps, and channels.
3. new allocates zeroed memory, while make allocates memory and initializes it.

Conclusion
In Golang, new and make are two keywords that are used for allocating memory.
new is used for creating a new instance of a variable, and it returns a pointer to the memory allocated.
make is used for creating slices, maps, and channels, and it returns the value of the type being created.
*/
