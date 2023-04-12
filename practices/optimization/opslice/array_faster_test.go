package opslice

import (
	"testing"
)

var opStr = [5]string{"one", "two", "3", "4", "five"}

func BenchmarkReturnSlice(b *testing.B) {
	var allSlice []string
	for sum := 0; sum < 100; sum++ {
		for _, v := range opStr {
			allSlice = append(allSlice, v)
		}
	}
}

func BenchmarkReturnArray(b *testing.B) {
	var allSlice [500]string
	for sum := 0; sum < 100; sum++ {
		for i, v := range opStr {
			allSlice[i] = v
		}
	}
}
