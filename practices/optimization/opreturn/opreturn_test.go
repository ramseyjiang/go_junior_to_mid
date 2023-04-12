package opreturn

import (
	"testing"
)

var s = [5]string{"one", "two", "3", "4", "five"}

func declareReturnTypeOnly() string {
	var allStr string
	for _, str := range s {
		allStr += str
	}

	return allStr
}

func declareReturnNameType() (allStr string) {
	for _, str := range s {
		allStr += str
	}

	return
}

func BenchmarkReturnTypeOnly(b *testing.B) {
	declareReturnTypeOnly()
}

func BenchmarkReturnNameTypeWithFmt(b *testing.B) {
	declareReturnNameType()
}
