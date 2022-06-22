package main

import "testing"

type U5 struct {
	a, b, c, d, e int
	h, i, j, k, l string
	o, p, q, r, s float32
}
type U4 struct {
	a, b, c, d int
	h, i, j, k string
	o, p, q, r float32
}
type U3 struct {
	b, c, d int
	h, i, j string
	o, p, q float32
}
type U2 struct {
	c, d int
}
type U1 struct {
	d int
}

func BenchmarkLoopArrayU5(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U5
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr); j++ {
			arr[j].d = j
		}
	}
}
func BenchmarkLoopArrayU4(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U4
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr); j++ {
			arr[j].d = j
		}
	}
}
func BenchmarkLoopArrayU3(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U3
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr); j++ {
			arr[j].d = j
		}
	}
}
func BenchmarkLoopArrayU2(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U2
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr); j++ {
			arr[j].d = j
		}
	}
}

func BenchmarkLoopArrayU1(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U1
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr); j++ {
			arr[j].d = j
		}
	}
}

func BenchmarkRangeArrayU5(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U5
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j].d = j
			_ = v
		}
	}
}
func BenchmarkRangeArrayU4(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U4
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j].d = j
			_ = v
		}
	}
}

func BenchmarkRangeArrayU3(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U3
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j].d = j
			_ = v
		}
	}
}
func BenchmarkRangeArrayU2(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U2
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j].d = j
			_ = v
		}
	}
}
func BenchmarkRangeArrayU1(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U1
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j].d = j
			_ = v
		}
	}
}
