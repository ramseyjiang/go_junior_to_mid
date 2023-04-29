package opslicemap

import "testing"

func VariableLenNoInitialSlice(size int) []int {
	slice := make([]int, 0)
	for i := 0; i < size; i++ {
		slice = append(slice, i)
	}
	return slice
}

func VariableLenShorterSlice(size int) []int {
	slice := make([]int, 100)
	for i := 0; i < size; i++ {
		slice = append(slice, i)
	}
	return slice
}

func VariableLenLargerSlice(size int) []int {
	slice := make([]int, 0, size)
	for i := 0; i < size; i++ {
		slice = append(slice, i)
	}
	return slice
}

func BenchmarkVariableLenNoInitialSlice(b *testing.B) {
	size := 10000
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		VariableLenNoInitialSlice(size)
	}
}

func BenchmarkVariableLenShorterSlice(b *testing.B) {
	size := 10000
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		VariableLenShorterSlice(size)
	}
}

func BenchmarkVariableLenLargerSlice(b *testing.B) {
	size := 10000
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		VariableLenLargerSlice(size)
	}
}
