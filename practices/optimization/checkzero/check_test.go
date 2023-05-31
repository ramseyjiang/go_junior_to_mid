package checkzero

import (
	"fmt"
	"testing"
)

func TestIsNotEmpty(t *testing.T) {
	var tests = []struct {
		input interface{}
		want  bool
	}{
		{"", false},
		{0, false},
		{false, false},
		{[3]int{0, 0, 0}, false},
		{map[string]int{}, false},
		{0, false},
		{"hello", true},
		{42, true},
		{nil, false},            // nil interface
		{(*int)(nil), false},    // nil pointer
		{(func())(nil), false},  // nil function
		{new(int), true},        // non-nil pointer
		{func() {}, true},       // non-nil function
		{[]int{}, false},        // empty slice
		{[]int{0, 0, 0}, true},  // non-empty slice
		{make(chan int), false}, // empty channel
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.input)
		t.Run(testname, func(t *testing.T) {
			flag := isNotEmpty(tt.input)
			if flag != tt.want {
				fmt.Println(tt.input)
				t.Errorf("got %v, want %v", flag, tt.want)
			}
		})
	}
}
