package main

import "testing"

type U5 struct {
	a, b, c, d, e int
}
type U4 struct {
	a, b, c, d int
}
type U3 struct {
	b, c, d int
}
type U2 struct {
	c, d int
}
type U1 struct {
	d int
}

func BenchmarkClassicForLoopLargeStructArrayU5(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U5
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr)-1; j++ {
			arr[j].d = j
		}
	}
}
func BenchmarkClassicForLoopLargeStructArrayU4(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U4
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr)-1; j++ {
			arr[j].d = j
		}
	}
}
func BenchmarkClassicForLoopLargeStructArrayU3(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U3
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr)-1; j++ {
			arr[j].d = j
		}
	}
}
func BenchmarkClassicForLoopLargeStructArrayU2(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U2
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr)-1; j++ {
			arr[j].d = j
		}
	}
}

func BenchmarkClassicForLoopLargeStructArrayU1(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U1
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr)-1; j++ {
			arr[j].d = j
		}
	}
}

func BenchmarkForRangeLargeStructArrayU5(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U5
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j].d = j
			_ = v
		}
	}
}
func BenchmarkForRangeLargeStructArrayU4(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U4
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j].d = j
			_ = v
		}
	}
}

func BenchmarkForRangeLargeStructArrayU3(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U3
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j].d = j
			_ = v
		}
	}
}
func BenchmarkForRangeLargeStructArrayU2(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U2
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j].d = j
			_ = v
		}
	}
}
func BenchmarkForRangeLargeStructArrayU1(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U1
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j].d = j
			_ = v
		}
	}
}

// no matter what type of structure, the performance of classic for loop traversal is the same,
// but the traversal performance of for range will decrease as the number of structure fields increases.
