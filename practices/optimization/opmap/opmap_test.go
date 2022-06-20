package opmap

import (
	"testing"
)

func BenchmarkDefineMapFirstWay(b *testing.B) {
	DefineMapFirstWay()
}

func BenchmarkDefineMapSecWay(b *testing.B) {
	DefineMapSecWay()
}

func BenchmarkDefineMapThirdWay(b *testing.B) {
	DefineMapThirdWay()
}

func BenchmarkDefineMapFourthWay(b *testing.B) {
	DefineMapFourthWay()
}
