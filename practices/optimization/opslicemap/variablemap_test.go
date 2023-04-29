package opslicemap

import "testing"

func NoCapacityMap(size int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < size; i++ {
		m[i] = i * 2
	}
	return m
}

func SmallerMap(size int) map[int]int {
	m := make(map[int]int, 100)
	for i := 0; i < size; i++ {
		m[i] = i * 2
	}
	return m
}

func LargerMap(size int) map[int]int {
	m := make(map[int]int, size)
	for i := 0; i < size; i++ {
		m[i] = i * 2
	}
	return m
}

func BenchmarkNoCapacityMap(b *testing.B) {
	size := 10000
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NoCapacityMap(size)
	}
}

func BenchmarkSmallerMap(b *testing.B) {
	size := 10000
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		SmallerMap(size)
	}
}

func BenchmarkLargerMap(b *testing.B) {
	size := 10000
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		LargerMap(size)
	}
}
