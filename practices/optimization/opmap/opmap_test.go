package opmap

import (
	"testing"
)

func BenchmarkDefineMapFirstWay(b *testing.B) {
	words := make(map[int]interface{})
	words[1] = "7"
	words[2] = "simple"
	words[3] = "ways"
	words[4] = "to"
	words[5] = "optimise"
	words[6] = "Golang"
}

func BenchmarkDefineMapSecWay(b *testing.B) {
	type str struct {
		word string
	}
	var words = map[int]str{}
	words[1] = str{"7"}
	words[2] = str{"simple"}
	words[3] = str{"ways"}
	words[4] = str{"to"}
	words[5] = str{"optimise"}
	words[6] = str{"Golang"}
}

func BenchmarkDefineMapThirdWay(b *testing.B) {
	words := make(map[int]string)
	words[1] = "7"
	words[2] = "simple"
	words[3] = "ways"
	words[4] = "to"
	words[5] = "optimise"
	words[6] = "Golang"
}

func BenchmarkDefineMapFourthWay(b *testing.B) {
	words := map[int]string{1: "7", 2: "simple", 3: "ways", 4: "to", 5: "optimise", 6: "Golang"}
	_ = words
}
