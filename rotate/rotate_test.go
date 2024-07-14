package rotate

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
		})
	}
}

func Benchmark_rotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	}
}

func Test_rotate2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate2(tt.args.nums, tt.args.k)
		})
	}
}

func Benchmark_rotate2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotate2([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	}
}

func Test_rotate3(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate3(tt.args.nums, tt.args.k)
		})
	}
}

func Benchmark_rotate3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotate3([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	}
}

func Test_rotate4(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate4(tt.args.nums, tt.args.k)
		})
	}
}

func Benchmark_rotate4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotate4([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	}
}
