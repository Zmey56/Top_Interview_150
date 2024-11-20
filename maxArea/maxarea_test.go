package maxArea

import (
	"math/rand"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Example 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "Example 2",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "Example 3",
			height: []int{4, 3, 2, 1, 4},
			want:   16,
		},
		{
			name:   "Example 4",
			height: []int{1, 2, 1},
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxArea(b *testing.B) {
	rng := rand.New(rand.NewSource(42))

	n := 100000
	height := make([]int, n)
	for i := 0; i < n; i++ {
		height[i] = rng.Intn(100000)
	}

	b.Run("FirstAlgorithm", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			maxArea(height)
		}
	})

	b.Run("SecondAlgorithm", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			maxAreaVer2(height)
		}
	})

	b.Run("ThirdAlgorithm", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			maxAreaVer3(height)
		}
	})
}
