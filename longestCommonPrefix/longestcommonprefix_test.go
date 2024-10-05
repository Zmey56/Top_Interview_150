package longestCommonPrefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "test 1",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "test 2",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "test 3",
			strs: []string{"", "racecar", "car"},
			want: "",
		},
		{
			name: "test 4",
			strs: []string{"", "racecar", ""},
			want: "",
		},
		{
			name: "test 5",
			strs: []string{"", "", ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefixVer4(tt.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_longestCommonPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestCommonPrefix([]string{"flower", "flow", "flight"})
	}
}

func Benchmark_longestCommonPrefixVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestCommonPrefixVer2([]string{"flower", "flow", "flight"})
	}
}

func Benchmark_longestCommonPrefixVer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestCommonPrefixVer3([]string{"flower", "flow", "flight"})
	}
}

func Benchmark_longestCommonPrefixVer4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestCommonPrefixVer3([]string{"flower", "flow", "flight"})
	}
}
