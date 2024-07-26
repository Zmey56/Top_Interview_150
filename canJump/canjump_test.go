package canJump

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{3, 2, 1, 0, 4},
			},
			want: false,
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{0},
			},
			want: true,
		},
		{
			name: "Test 4",
			args: args{
				nums: []int{1, 0},
			},
			want: true,
		},
		{
			name: "Test 5",
			args: args{
				nums: []int{1, 1, 1, 0},
			},
			want: true,
		},
		{
			name: "Test 6",
			args: args{
				nums: []int{1, 1, 1, 0, 1},
			},
			want: false,
		},
		{
			name: "Test 7",
			args: args{
				nums: []int{3, 0, 8, 2, 0, 0, 1},
			},
			want: true,
		},
		{
			name: "Test 8",
			args: args{
				nums: []int{0, 2, 3},
			},
			want: false,
		},
		{
			name: "Test 9",
			args: args{
				nums: []int{1, 0, 1, 0},
			},
			want: false,
		},
		{
			name: "Test 10",
			args: args{
				nums: []int{2, 1, 2, 2, 1, 2, 2, 2},
			},
			want: true,
		},
		{
			name: "Test 11",
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_canJump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canJump([]int{2, 3, 1, 1, 4})
	}
}

func Benchmark_canJump2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canJump2([]int{2, 3, 1, 1, 4})
	}
}

func Benchmark_canJump3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canJump3([]int{2, 3, 1, 1, 4})
	}
}
