package main

import (
	"fmt"
)

/*
	The only difference between a slice and an array is the missing length between the brackets.
	Slices are always associated with some array, and although they can never be longer than the array, they can be smaller.

	One way to create a slice,
	y := make([]float64, 5)

	The make function also allows a third parameter: x := make([]float64, 5, 10)

	Another way to create slices is to use the [low : high] expression:
	arr := [3]float64{1,2,3}
	x := arr[0:2]

	low is the index of where to start the slice and high is the index of where to end it,
	but not including the index itself.

	For convenience, we are also allowed to omit low, high, or even both low and high.
	arr[0:] is the same as arr[0:len(arr)], arr[:5] is the same as arr[0:5], and
	arr[:] is the same as arr[0:len(arr)].
*/
func main() {
	arr := [3]float64{1, 2, 3}
	x := arr[0:2]  // It means x includes the first 5 elements in arr.
	fmt.Println(x) // [1 2]

	slice := appendSlice(x)
	fmt.Println("append alice results: ", slice)

	slice1, slice2, slice3 := copySlice()
	fmt.Println(slice1, slice2, slice3)
}

/*
	Append adds elements onto the end of a slice.
	If there’s sufficient capacity in the underlying array,
	the element is placed after the last element and the length is incremented.
	However, if there is not enough capacity, a new array is created, all of existing elements are
*/
func appendSlice(slice []float64) []float64 {
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 4, 5)

	return slice
}

/*
	copy takes two arguments: dst and src. All entries in src are copied into dst overwriting whatever is there.
	If the lengths of the two slices are not the same, the smaller of the two will be used.
*/
func copySlice() ([]int, []int, int) {
	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 3) // It means slice2 maximum length is 2.

	// copy() returns the number of elements copied, which will be the minimum of len(src) and len(dst).
	slice3 := copy(slice2, slice1)
	fmt.Println(slice1, slice2, slice3)

	return slice1, slice2, slice3
}
