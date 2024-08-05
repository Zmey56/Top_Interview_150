package hIndex

import "testing"

func Test_hIndex(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		want      int
	}{
		{
			name:      "Test case 1",
			citations: []int{3, 0, 6, 1, 5},
			want:      3,
		},
		{
			name:      "Test case 2",
			citations: []int{1, 3, 1},
			want:      1,
		},
		{
			name:      "Test case 3",
			citations: []int{0, 1, 3, 5, 6},
			want:      3,
		},
		{
			name:      "Test case 4",
			citations: []int{0, 1, 2, 3, 4, 5, 6},
			want:      3,
		},
		{
			name:      "Test case 5",
			citations: []int{0, 1, 2, 3, 4, 5, 6, 7},
			want:      4,
		},
		{
			name:      "Test case 6",
			citations: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			want:      4,
		},
		{
			name:      "Test case 7",
			citations: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want:      5,
		},
		{
			name:      "Test case 8",
			citations: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:      5,
		},
		{
			name:      "Test case 9",
			citations: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			want:      6,
		},
		{
			name:      "Test case 10",
			citations: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			want:      6,
		},
		{
			name:      "Test case 11",
			citations: []int{1, 1},
			want:      1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex2(tt.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Benchmark_hIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hIndex([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	}
}

func Benchmark_hIndex2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hIndex2([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	}
}

func Benchmark_hIndex3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hIndex3([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	}
}
