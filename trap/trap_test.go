package trap

import (
	"fmt"
	"testing"
)

func Test_trap(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{3, 2, 1, 2, 1}, 1},
		{[]int{0, 1, 0, 2}, 1},
		{[]int{0, 1, 0, 2, 1, 0, 1, 3}, 5},
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{4, 2, 3}, 1},
		{[]int{4, 2, 3, 1}, 1},
		{[]int{4, 2, 3, 1, 2}, 2},
		{[]int{4, 2, 3, 1, 2, 1}, 2},
		{[]int{2, 3, 4, 1, 2, 1}, 1},
		{[]int{3, 2, 0, 2, 4, 1, 2}, 6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.height), func(t *testing.T) {
			if got := trap(tt.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_trap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	}
}

func Benchmark_trapVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trapVer2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	}
}

func Benchmark_trapVer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trapVer3([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	}
}
