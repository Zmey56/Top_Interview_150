package canJumpII

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{2, 3, 0, 1, 4},
			},
			want: 2,
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{2, 3, 0, 1, 4, 1, 1, 1, 1},
			},
			want: 3,
		},
		{
			name: "Test 4",
			args: args{
				nums: []int{2, 3, 0, 1, 4, 1, 1, 1, 1, 1},
			},
			want: 4,
		},
		{
			name: "Test 5",
			args: args{
				nums: []int{3, 4, 3, 1, 0, 7, 0, 3, 0, 2, 0, 3},
			},
			want: 3,
		},
		{
			name: "Test 6",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump2(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_jump(b *testing.B) {
	nums := []int{3, 4, 3, 1, 0, 7, 0, 3, 0, 2, 0, 3}
	for i := 0; i < b.N; i++ {
		jump(nums)
	}
}

func Benchmark_jump2(b *testing.B) {
	nums := []int{3, 4, 3, 1, 0, 7, 0, 3, 0, 2, 0, 3}
	for i := 0; i < b.N; i++ {
		jump2(nums)
	}
}
