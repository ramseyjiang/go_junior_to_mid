package main

import (
	"fmt"
	"strconv" // if does not import strconv, when using strconv.Itoa(i), it will have an error
)

func main() {
	i := 24
	fmt.Printf("%v, %T\n", i, i) // 24, int

	practiceStrToInt(i)

	practiceIntToStr()
}

// int/int64 convert to string, Itoa base on strconv.FormatInt(i, 10)
func practiceStrToInt(i int) {
	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j) // 24, string

	var n int64 = 97
	s := strconv.FormatInt(n, 16)
	fmt.Printf("%v, %T\n", s, s) // 61, string, hexadecimal, it means base 16 for "0x"

	var m int64 = 97
	t := strconv.FormatInt(m, 10)
	fmt.Printf("%v, %T\n", t, t) // 97, string, decimal
}

// string convert to int/int64
func practiceIntToStr() {
	k := "97"

	n1, err1 := strconv.Atoi(k) // using strconv.Atoi to convert a string to int decimal
	if err1 == nil {
		fmt.Printf("%d, %T\n", n1, n1) // 97, int
	} else {
		fmt.Println(k, "is not an integer.")
	}

	s := "97"
	n2, err2 := strconv.ParseInt(s, 10, 64)
	if err2 == nil {
		fmt.Printf("%d, %T", n2, n2) // 97, int64
	} else {
		fmt.Println(s, "is not an integer.")
	}
}
