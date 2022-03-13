package main

import (
	"fmt"
	"time"
)

func main() {
	x := arrayDefine()
	fmt.Println("define an int array and return it", x) // [0 0 0 0 100]

	floatArrayCalculate() // The first way is used to calculate a float array, not highly recommend.
	totalFloat64 := forRangeArrayLoop()
	fmt.Println("Highly recommend this way to calculate a float array", totalFloat64)

	date1, dateArr := timeValueArray()
	fmt.Println("define an time.Time array and return it", date1, dateArr[0])

	notes, primes := returnMultipleArrayTypes()
	fmt.Println(notes)          // [do re mi]
	fmt.Printf("%#v\n", notes)  // [3]string{"do", "re", "mi"}
	fmt.Println(primes)         // [2 3 5 7 11]
	fmt.Printf("%#v\n", primes) // [5]int{2, 3, 5, 7, 11}

	arrayCompare() // An example of array compare, the value, type, and element order must all the same.

	a, b := multiDimensionalArray() // When array elements of an array are arrays, then it’s called a multi-dimensional array.
	fmt.Println(a, b, b[0])         // Please observe the different array length when they are defined.
}

func arrayDefine() [5]int {
	var x [5]int
	x[4] = 100

	return x
}

/*
	define an array   	var myArray [4] string
	4 means number of elements array will hold. string means type of elements array will hold.
	Elements in an array are numbered, starting with 0. An element's number is called its index.

	A single underscore (_) is used to tell the compiler that we don’t need this. Because in go, it doesn't allow useless variable exists.
	value is the same as x[i].
	The following is similar as php foreach $array as $key => $value.
	for _, value := range x {
		total += value
	}

	Using for...range, it will make loop an array safely.
*/
func forRangeArrayLoop() float64 {
	/*
		It is always an array type at the beginning.
		And then with the number of elements it will hold in square brackets, followed by the type of its elements.
		After that, it will be followed by a list in curly braces and the initial values each element should have.
	*/
	x := [5]float64{98, 93, 77, 82, 83} // This way defines an array called Array literals
	var total float64

	// _ it will hold the array index, if you want to use index, then using index to replace the _.
	for _, value := range x {
		total += value
	}

	return total / float64(len(x))
}

func floatArrayCalculate() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)
}

func timeValueArray() (time.Time, [3]time.Time) {
	var dates [3]time.Time // Create an array of three Time values. In go, Time is a type of value.

	dates[0] = time.Unix(1257894000, 0) // Assign a value to an element.
	dates[1] = time.Unix(1447920000, 0)
	dates[2] = time.Unix(1508632200, 0)

	return dates[1], dates
}

// Using "%#v" in Printf() , it will formats values as they'd appear in Go code.
func returnMultipleArrayTypes() ([3]string, [5]int) {
	var notes = [3]string{"do", "re", "mi"}
	primes := [5]int{2, 3, 5, 7, 11}

	return notes, primes
}

// For an array to be the equal or the same as the second array, both array should be of the same type, must have
// the same elements in them and all elements must be in the same order.
func arrayCompare() {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 3, 2}
	c := [3]int{1, 2, 1}
	d := [3]int{1, 2, 3}

	fmt.Println("a == b ", a == b)
	fmt.Println("a == c ", a == c)
	fmt.Println("a == d ", a == d)
}

// [[1 2] [3 4] [0 0]]
// When define an array using "...", the array's length is flexible.
// The array inside should have the same length
// You define a dynamic array, but when you return an array it must be a stable length.
// Otherwise, you can use interface to define a type by yourself.
func multiDimensionalArray() ([3][2]int, [3][2]int) {
	a := [3][2]int{
		{1, 2},
		{3, 4},
	}

	b := [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	return a, b
}
