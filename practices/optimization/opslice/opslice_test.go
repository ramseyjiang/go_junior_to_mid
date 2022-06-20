package opslice

import (
	"testing"
)

func BenchmarkReturnSliceNoLength(b *testing.B) {
	_ = declareReturnSliceNoLength(opStr)
}

func BenchmarkReturnSliceWithLength(b *testing.B) {
	_ = declareReturnSliceWithLength(opStr)
}
