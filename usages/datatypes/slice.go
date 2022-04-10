package main

import (
	"fmt"
	"time"
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
	arr := [3]int64{1, 2, 3}
	x := arr[0:2]  // It means x includes the first 5 elements in arr.
	fmt.Println(x) // [1 2]

	slice := appendSlice(x, arr)
	fmt.Println("append alice results: ", slice) // [1 2 7 4 5]

	slice1, slice2, slice3 := copySlice()
	fmt.Println(slice1, slice2, slice3) // [1 2 3 4] [1 2 3] 3

	trickSlice()

	efficiencySlice()
}

/*
	Append adds elements onto the end of a slice.
	If there’s sufficient capacity in the underlying array,
	the element is placed after the last element and the length is incremented.
	However, if there is not enough capacity, a new array is created, all of existing elements are
*/
func appendSlice(slice []int64, arr [3]int64) []int64 {
	arr1 := make([]int64, 0, len(arr))
	arr2 := make([]int64, len(arr))

	slice = append(slice, 7)
	slice = append(slice, 4, 5)

	slice1 := append(arr1, slice...)
	fmt.Println("arr1 append result is ", slice1)
	slice2 := append(arr2, slice...)
	fmt.Println("arr2 append result is ", slice2)

	return slice
}

/*
	copy takes two arguments: dst and src. All entries in src are copied into dst overwriting whatever is there.
	If the lengths of the two slices are not the same, the smaller of the two will be used.
*/
func copySlice() ([]int, []int, int) {
	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 3) // It means slice2 maximum length is 3.

	// copy() returns the number of elements copied, which will be the minimum of len(src) and len(dst).
	slice3 := copy(slice2, slice1)
	fmt.Println(slice1, slice2, slice3)

	return slice1, slice2, slice3
}

func trickSlice() {
	numbers := []int{1, 3, 7, 11}
	newNumbersA := numbers
	newNumbersB := numbers[1:4]

	numbers[2] = 10
	fmt.Println(newNumbersA, numbers, newNumbersB)

	newNumbersA[3] = 50
	fmt.Println(newNumbersA, numbers, newNumbersB)

	newNumbersB[1] = 30
	fmt.Println(newNumbersA, numbers, newNumbersB)

	// On the above, newNumbersA and newNumbersB are using the same original array, that's why update one,
	// then others results are changed.
	// newNumbersC is defined using append, and it appends to a new slice, so it is changed, it won't change the original slice.
	newNumbersC := append([]int{}, numbers...)
	fmt.Println(newNumbersC)
	newNumbersB[1] = 20
	fmt.Println(newNumbersA, numbers, newNumbersB, newNumbersC)
}

func efficiencySlice() {
	start := time.Now()
	inputSlice := [6]int{0, 1, 2, 3, 4, 5}

	var sum int
	var firstSlice []int
	for sum = 0; sum < 100000; sum++ {
		for _, v := range inputSlice {
			firstSlice = append(firstSlice, v)
		}
	}
	elapsed1 := time.Since(start)
	fmt.Println(elapsed1, len(firstSlice))

	start = time.Now()
	secondSlice := make([]int, len(firstSlice))
	for sum = 0; sum < 100000; sum++ {
		for i, v := range inputSlice {
			secondSlice[i] = v
		}
	}
	elapsed2 := time.Since(start)
	fmt.Println(elapsed2, len(secondSlice))

	start = time.Now()
	thirdSlice := make([]int, 0, len(firstSlice))
	for sum = 0; sum < 100000; sum++ {
		for _, v := range inputSlice {
			thirdSlice = append(thirdSlice, v)
		}
	}
	elapsed3 := time.Since(start)
	fmt.Println(elapsed3, len(secondSlice))
}
