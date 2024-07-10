package majorityElement

import "testing"

func Test_majorityElement(t *testing.T) {
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
				nums: []int{3, 2, 3},
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_majorityElement(b *testing.B) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	for i := 0; i < b.N; i++ {
		majorityElement(nums)
	}
}

func Test_majorityElementVer2(t *testing.T) {
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
				nums: []int{3, 2, 3},
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElementVer2(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_majorityElementVer2(b *testing.B) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	for i := 0; i < b.N; i++ {
		majorityElementVer2(nums)
	}
}

func Test_majorityElementVer3(t *testing.T) {
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
				nums: []int{3, 2, 3},
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElementVer2(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_majorityElementVer3(b *testing.B) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	for i := 0; i < b.N; i++ {
		majorityElementVer3(nums)
	}
}

func Test_majorityElementVer4(t *testing.T) {
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
				nums: []int{3, 2, 3},
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElementVer2(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_majorityElementVer4(b *testing.B) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	for i := 0; i < b.N; i++ {
		majorityElementVer4(nums)
	}
}

func Test_majorityElementVer5(t *testing.T) {
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
				nums: []int{3, 2, 3},
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElementVer2(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_majorityElementVer5(b *testing.B) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	for i := 0; i < b.N; i++ {
		majorityElementVer5(nums)
	}
}

func Test_majorityElementVer6(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{3, 2, 3},
			},
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			majorityElementVer6(tt.args.nums)
		})
	}
}

func Benchmark_majorityElementVer6(b *testing.B) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	for i := 0; i < b.N; i++ {
		majorityElementVer6(nums)
	}
}
