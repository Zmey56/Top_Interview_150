package canCompleteCircuit

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	tests := []struct {
		name string
		gas  []int
		cost []int
		want int
	}{
		{
			name: "Example 1",
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
			want: 3,
		},
		{
			name: "Example 2",
			gas:  []int{2, 3, 4},
			cost: []int{3, 4, 3},
			want: -1,
		},
		{
			name: "Example 3",
			gas:  []int{5, 1, 2, 3, 4},
			cost: []int{4, 4, 1, 5, 1},
			want: 4,
		},
		{
			name: "Example 4",
			gas:  []int{2, 3, 4},
			cost: []int{3, 4, 5},
			want: -1,
		},
		{
			name: "Example 5",
			gas:  []int{4, 5, 2, 6, 5, 3},
			cost: []int{3, 2, 7, 3, 2, 9},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteCircuitConcurrent(tt.gas, tt.cost); got != tt.want {
				t.Errorf("canCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_canCompleteCircuit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	}
}

func Benchmark_canCompleteCircuitVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canCompleteCircuitVer2([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	}
}

func Benchmark_canCompleteCircuitConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canCompleteCircuitConcurrent([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	}
}
