package opslicemap

import "testing"

func FixedLenSlice(size int) []int {
	slice := make([]int, 0, size)
	for i := 0; i < size; i++ {
		slice = append(slice, i)
	}
	return slice
}

func BenchmarkFixedLenSliceShort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FixedLenSlice(100)
	}
}

func BenchmarkFixedLengthSliceLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FixedLenSlice(10000)
	}
}
