package opinterface

import "testing"

// Fixed type approach
func DoSomethingFixed(value int) int {
	return value
}

// Interface approach
func DoSomethingInterface(value interface{}) interface{} {
	return value
}

// Generics approach (Go 1.18+ required)
func DoSomethingGeneric[T any](value T) T {
	return value
}

func BenchmarkDoSomethingGeneric(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DoSomethingGeneric(66)
	}
}

func BenchmarkDoSomethingInterface(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DoSomethingInterface(66)
	}
}

func BenchmarkDoSomethingFixed(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DoSomethingFixed(66)
	}
}
