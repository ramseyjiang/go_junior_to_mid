package main

import (
	"fmt"
	"time"
	"unsafe"
)

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

func main() {
	fmt.Printf("Size of %T struct: %d bytes\n", employee1, unsafe.Sizeof(employee1))

	fmt.Printf("Size of %T struct: %d bytes\n", employee2, unsafe.Sizeof(employee2))

	fmt.Printf("Size of %T struct: %d bytes\n", employee3, unsafe.Sizeof(employee3))
}
