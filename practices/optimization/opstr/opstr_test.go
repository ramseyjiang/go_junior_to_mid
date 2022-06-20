package opstr

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkWithSprintfSmallString(b *testing.B) {
	data := generateRandomLengthStrings(smallString)
	var s string
	for n := 0; n < b.N; n++ {
		s = fmt.Sprintf(s, data)
		_ = s
	}
}
func BenchmarkWithSprintfLongerString(b *testing.B) {
	data := generateRandomLengthStrings(longString)
	var s string
	for n := 0; n < b.N; n++ {
		s = fmt.Sprintf(s, data)
		_ = s
	}
}

func BenchmarkWithOperatorSmallString(b *testing.B) {
	data := generateRandomLengthStrings(smallString)
	str := strings.Join(data, "")
	var s string
	for n := 0; n < b.N; n++ {
		s += str
		_ = s
	}
}

func BenchmarkWithOperatorLongerString(b *testing.B) {
	data := generateRandomLengthStrings(longString)
	str := strings.Join(data, "")
	var s string
	for n := 0; n < b.N; n++ {
		s += str
		_ = s
	}
}

func BenchmarkWithJoinSmallerString(b *testing.B) {
	data := generateRandomLengthStrings(smallString)
	var s string
	for n := 0; n < b.N; n++ {
		s = strings.Join(data, " ")
		_ = s
	}
}
func BenchmarkWithJoinLongerString(b *testing.B) {
	data := generateRandomLengthStrings(longString)
	var s string
	for n := 0; n < b.N; n++ {
		s = strings.Join(data, " ")
		_ = s
	}
}

func BenchmarkWithStringBuilderSmallerString(b *testing.B) {
	data := generateRandomLengthStrings(smallString)
	var builder strings.Builder
	var s string

	for n := 0; n < b.N; n++ {
		for _, s := range data {
			builder.WriteString(s)
		}
		s = builder.String()
		_ = s
	}
}
func BenchmarkWithStringBuilderLongerString(b *testing.B) {
	data := generateRandomLengthStrings(longString)
	var builder strings.Builder
	var s string

	for n := 0; n < b.N; n++ {
		for _, s := range data {
			builder.WriteString(s)
		}
		s = builder.String()
		_ = s
	}
}
