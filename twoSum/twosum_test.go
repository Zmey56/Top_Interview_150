package twoSum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Example 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{1, 2},
		},
		{
			name:   "Example 2",
			nums:   []int{2, 3, 4},
			target: 6,
			want:   []int{1, 3},
		},
		{
			name:   "Example 3",
			nums:   []int{-1, 0},
			target: -1,
			want:   []int{1, 2},
		},
		{
			name: "Example 4",
			nums: []int{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
				21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
				31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
				41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
				51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
				61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
				71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
				81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
				91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
			},
			target: 150,
			want:   []int{50, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_twoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{2, 7, 11, 15}, 9)
	}
}

func Benchmark_twoSumVeer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSumVer2([]int{2, 7, 11, 15}, 9)
	}
}

func Benchmark_twoSumVer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSumVer3([]int{2, 7, 11, 15}, 9)
	}
}

func BenchmarkTwoSum(b *testing.B) {
	numbers := make([]int, 30000)
	for i := 0; i < 30000; i++ {
		numbers[i] = i + 1
	}
	target := 45000

	b.Run("FirstAlgorithm", func(b *testing.B) {
		var result []int
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = twoSum(numbers, target)
		}
		if len(result) == 0 {
			b.Fatal("No result found")
		}
	})

	b.Run("SecondAlgorithm", func(b *testing.B) {
		var result []int
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = twoSumVer2(numbers, target)
		}
		if len(result) == 0 {
			b.Fatal("No result found")
		}
	})

	b.Run("ThirdAlgorithm", func(b *testing.B) {
		var result []int
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = twoSumVer3(numbers, target)
		}
		if len(result) == 0 {
			b.Fatal("No result found")
		}
	})
}
