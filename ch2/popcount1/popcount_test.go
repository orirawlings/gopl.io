package popcount

import (
	"fmt"
	"testing"
)

func ExamplePopCount() {
	fmt.Println(PopCount(0x123456789ABCDEF))
	// Output: 32
}

func ExamplePopCountShift() {
	fmt.Println(PopCountShift(0x123456789ABCDEF))
	// Output: 32
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(0x1234567890ABCDEF)
	}
}
