package isSubsequence

import "testing"

func Test_isSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "Example 2",
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.s, tt.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_isSubsequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isSubsequence("abc", "ahbgdc")
	}
}

func Benchmark_isSubsequenceVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isSubsequenceVer2("abc", "ahbgdc")
	}
}

func Benchmark_isSubsequenceVer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isSubsequenceVer3("abc", "ahbgdc")
	}
}
