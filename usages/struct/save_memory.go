package main

import (
	"fmt"
	"unsafe"
)

/**
We know that CPU 32-bit and CPU 64-bit are the most popular types of CPU.
Imagine that we have a CPU 64-bit, the CPU has transfer capability 64-bit data per clock cycle.
The clock cycle is how much time needed for CPU to process an information.

For example, when you go to a supermarket, you will have a trolley. The maximum capacity of the trolley
is the same as a cpu maximum process ability of each clock cycle.

In golang, data type and size list is following:
Data Type 			Size
bool				1 byte
int16				2 byte
int32				4 byte
int64				8 byte
string				16 byte
float32				4 byte
*/

type Employee1 struct {
	IsActive  bool
	Age       int64
	IsMarried bool
	Name      string
	weight    int32
	height    int16
	Photo     float32
}

type Employee2 struct {
	Name      string
	Age       int64
	Photo     float32
	weight    int32
	height    int16
	IsMarried bool
	IsActive  bool
}

type Employee3 struct {
	Name      string
	IsActive  bool
	Age       int64
	IsMarried bool
	weight    int32
	height    int16
	Photo     float32
}

var employee1 Employee1
var employee2 Employee2
var employee3 Employee3

/**
So, the order is important, but the key point is the space of each element occupy.
*/
func main() {
	fmt.Printf("Size of %T struct: %d bytes\n", employee1, unsafe.Sizeof(employee1))
	fmt.Printf("Size of %T struct: %d bytes\n", employee2, unsafe.Sizeof(employee2))
	fmt.Printf("Size of %T struct: %d bytes\n", employee3, unsafe.Sizeof(employee3))
}
