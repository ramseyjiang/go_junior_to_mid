package main

import (
	"fmt"
	"time"
	"unsafe"
)

/**
CPU 32-bit and CPU 64-bit are the most popular types of CPU.
Imagine that we have a CPU 64-bit, the CPU has transfer capability 64-bit data per clock cycle.
The clock cycle is how much time needed for CPU to process an information.

For example, when you go to a supermarket, you will have a trolley. The maximum capacity of the trolley
is the same as a cpu maximum process ability of each clock cycle.

In golang, data type and size list is following:
Data Type 			Size
bool				1 byte
int16				2 bytes
int32				4 bytes
int64				8 bytes
int					8 bytes
string				16 bytes
float32				4 bytes
float64				8 bytes
uint32				4 bytes
uint64				8 bytes
nil interface{}		16 bytes
time.Time			24 bytes   // it is a struct, so its size is unstable.
time.Timer			80 bytes   // it is a struct, so its size is unstable.
time.Duration		8 bytes
[]byte				24 bytes
*/

type commonStruct struct {
	createAt time.Time
	updateAt time.Timer
	timeout  time.Duration
	jsonStr  []byte
}

type Employee1 struct {
	IsActive  bool
	Age       int64
	IsMarried bool
	Name      string
	weight    int32
	height    int16
	PhotoLen  float64
	PhotoWid  float32
	intNum    int
	length    int8
	common    commonStruct
}

type Employee2 struct {
	Name      string
	Age       int64
	intNum    int
	PhotoLen  float64
	PhotoWid  float32
	weight    int32
	height    int16
	length    int8
	IsMarried bool
	IsActive  bool
	common    commonStruct
}

type Employee3 struct {
	Name      string
	IsActive  bool
	intNum    int
	length    int8
	Age       int64
	IsMarried bool
	weight    int32
	height    int16
	PhotoWid  float32
	PhotoLen  float64
	common    commonStruct
}

var employee1 Employee1
var employee2 Employee2
var employee3 Employee3

/**
So, the order is important, but the key point is the space of each element occupy.
Use fmt.Println(unsafe.Sizeof(employee1.xxx)) to get size of variable, if unsafe.Sizeof() does not in fmt.xxx, it won't work.
*/
func main() {
	fmt.Printf("Size of %T struct: %d bytes\n", employee1, unsafe.Sizeof(employee1))

	fmt.Printf("Size of %T struct: %d bytes\n", employee2, unsafe.Sizeof(employee2))

	fmt.Printf("Size of %T struct: %d bytes\n", employee3, unsafe.Sizeof(employee3))
}
