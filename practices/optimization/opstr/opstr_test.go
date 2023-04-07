package opstr

import (
	"strings"
	"testing"
)

func BenchmarkSmallStringPlus(b *testing.B) {
	words := []string{"hello", "world", "this", "is", "a", "test"}
	str := strings.Join(words, "")
	var s string
	for n := 0; n < b.N; n++ {
		s += str
		_ = s
	}
}

func BenchmarkSmallStringsJoin(b *testing.B) {
	words := []string{"hello", "world", "this", "is", "a", "test"}

	for i := 0; i < b.N; i++ {
		result := strings.Join(words, " ")
		_ = result
	}
}

func BenchmarkSmallStringsBuilder(b *testing.B) {
	words := []string{"hello", "world", "this", "is", "a", "test"}

	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for _, word := range words {
			builder.WriteString(word)
			builder.WriteString(" ")
		}
		result := builder.String()
		_ = result
	}
}

func BenchmarkLongStringsJoin(b *testing.B) {
	words := make([]string, 1000)
	for i := range words {
		words[i] = "test"
	}

	for i := 0; i < b.N; i++ {
		result := strings.Join(words, " ")
		_ = result
	}
}

func BenchmarkLongStringsBuilder(b *testing.B) {
	words := make([]string, 1000)
	for i := range words {
		words[i] = "test"
	}

	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for _, word := range words {
			builder.WriteString(word)
			builder.WriteString(" ")
		}
		result := builder.String()
		_ = result
	}
}
