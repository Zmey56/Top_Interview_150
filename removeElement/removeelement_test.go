package removeElement

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_removeElement(b *testing.B) {
	nums := []int{3, 2, 2, 3}
	for i := 0; i < b.N; i++ {
		removeElement(nums, 3)
	}
}

func Test_removeElementVer2(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElementVer2(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElementVer2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_removeElementVer2(b *testing.B) {
	nums := []int{3, 2, 2, 3}
	for i := 0; i < b.N; i++ {
		removeElementVer2(nums, 3)
	}
}

func Test_removeElementVer3(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElementVer3(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElementVer2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_removeElementVer3(b *testing.B) {
	nums := []int{3, 2, 2, 3}
	for i := 0; i < b.N; i++ {
		removeElementVer3(nums, 3)
	}
}
