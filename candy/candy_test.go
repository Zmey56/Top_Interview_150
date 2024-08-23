package candy

import "testing"

func Test_candy(t *testing.T) {
	tests := []struct {
		name    string
		ratings []int
		want    int
	}{
		{
			name:    "Example 1",
			ratings: []int{1, 0, 2},
			want:    5,
		},
		{
			name:    "Example 2",
			ratings: []int{1, 2, 2},
			want:    4,
		},
		{
			name:    "Example 3",
			ratings: []int{1, 2, 87, 87, 87, 2, 1},
			want:    13,
		},
		{
			name:    "Example 4",
			ratings: []int{1, 3, 4, 5, 2},
			want:    11,
		},
		{
			name:    "Example 5",
			ratings: []int{1, 2, 3, 4, 5},
			want:    15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_candy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		candy([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func Benchmark_candyVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		candyVer2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func Benchmark_candyVer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		candyVer3([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func Benchmark_candyVer4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		candyVer4([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}
