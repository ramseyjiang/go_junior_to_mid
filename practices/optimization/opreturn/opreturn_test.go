package opreturn

import (
	"testing"
)

func BenchmarkReturnTypeOnly(b *testing.B) {
	s := declareReturnTypeOnly()
	_ = s
}

func BenchmarkReturnNameTypeWithFmt(b *testing.B) {
	s := declareReturnNameTypeWithFmt()
	_ = s
}

func BenchmarkReturnNameTypeWithLog(b *testing.B) {
	s := declareReturnNameTypeWithLog()
	_ = s
}
