package opslice

import (
	"testing"
)

const size = 1000

func processArray(arr [size]int, length int) int {
	return arr[length-1]
}

func processSlice(slc []int) int {
	return slc[len(slc)-1]
}

func BenchmarkArray(b *testing.B) {
	arr := [size]int{}
	length := 0

	for i := 0; i < b.N; i++ {
		length = (length + 1) % (size + 1)
		if length == 0 {
			continue
		}

		arr[length-1] = i
		_ = processArray(arr, length)
	}
}

func BenchmarkSlice(b *testing.B) {
	slc := make([]int, 0, size)

	for i := 0; i < b.N; i++ {
		if len(slc) >= size {
			slc = slc[1:]
		}
		slc = append(slc, i)
		_ = processSlice(slc)
	}
}
