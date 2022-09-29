package main

import (
	"testing"
)

const it = uint(1 << 24)

func BenchmarkSetWithBoolValueWrite(b *testing.B) {
	set := make(map[uint]bool)

	for i := uint(0); i < it; i++ {
		set[i] = true
	}
}

func BenchmarkSetWithStructValueWrite(b *testing.B) {
	set := make(map[uint]struct{})

	for i := uint(0); i < it; i++ {
		set[i] = struct{}{}
	}
}

func BenchmarkSetWithInterfaceValueWrite(b *testing.B) {
	set := make(map[uint]interface{})

	for i := uint(0); i < it; i++ {
		set[i] = struct{}{}
	}
}

/**
From the result below, map[]struct{} is 5% faster in time and 10% less memory consumption comparing to map[]bool when it comes to a big Set.
BenchmarkSetWithBoolValueWrite
BenchmarkSetWithBoolValueWrite-10                      1        2001878000 ns/op        884605976 B/op    613620 allocs/op
BenchmarkSetWithBoolValueWrite-10                      1        1961659750 ns/op        886793720 B/op    614112 allocs/op
BenchmarkSetWithBoolValueWrite-10                      1        1996489750 ns/op        884576056 B/op    613451 allocs/op
BenchmarkSetWithStructValueWrite
BenchmarkSetWithStructValueWrite-10                    1        1951849667 ns/op        803449584 B/op    613750 allocs/op
BenchmarkSetWithStructValueWrite-10                    1        1920334916 ns/op        803442296 B/op    613125 allocs/op
BenchmarkSetWithStructValueWrite-10                    1        1932203167 ns/op        805717072 B/op    614140 allocs/op
BenchmarkSetWithInterfaceValueWrite
BenchmarkSetWithInterfaceValueWrite-10                 1        2460922958 ns/op        1981824856 B/op   615051 allocs/op
BenchmarkSetWithInterfaceValueWrite-10                 1        2438848708 ns/op        1981418584 B/op   613095 allocs/op
BenchmarkSetWithInterfaceValueWrite-10                 1        2440771708 ns/op        1981532472 B/op   613642 allocs/op


Although map[]struct{} is faster and consumes less memory, it is more obvious when it comes to a big Set.

*/
