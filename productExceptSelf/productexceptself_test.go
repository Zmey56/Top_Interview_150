package productExceptSelf

import "testing"

func Test_productExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "Example 2",
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelfVer5(tt.nums); !equal(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Benchmark_productExceptSelf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		productExceptSelf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func Benchmark_productExceptSelfVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		productExceptSelfVer2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func Benchmark_productExceptSelfVer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		productExceptSelfVer3([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func Benchmark_productExceptSelfVer4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		productExceptSelfVer4([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func Benchmark_productExceptSelfVer5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		productExceptSelfVer5([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}
