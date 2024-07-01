package merge

import "testing"

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
		},
		{
			name: "Test 2",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
		},
		{
			name: "Test 3",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		})
	}
}

func Benchmark_merge(b *testing.B) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	for i := 0; i < b.N; i++ {
		merge(nums1, 3, nums2, 3)
	}

}

func Test_mergeVer2(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
		},
		{
			name: "Test 2",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
		},
		{
			name: "Test 3",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeVer2(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		})
	}
}

func Benchmark_mergeVer2(b *testing.B) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	for i := 0; i < b.N; i++ {
		mergeVer2(nums1, 3, nums2, 3)
	}

}
