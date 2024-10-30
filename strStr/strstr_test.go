package strStr

import "testing"

func Test_strStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{
			name:     "Example 1",
			haystack: "hello",
			needle:   "ll",
			want:     2,
		},
		{
			name:     "Example 2",
			haystack: "aaaaa",
			needle:   "bba",
			want:     -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.haystack, tt.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_strStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strStr("hello", "ll")
	}
}

func Benchmark_strStrVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strStrVer2("hello", "ll")
	}
}

func Benchmark_strStrVer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strStrVer3("hello", "ll")
	}
}
