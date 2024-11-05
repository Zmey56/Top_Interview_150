package isPalindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "Example 2",
			s:    "race a car",
			want: false,
		},
		{
			name: "Example 3",
			s:    "Able was I ere I saw Elba",
			want: true,
		},
		{
			name: "Example 4",
			s:    " ",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_isPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome("A man, a plan, a canal: Panama")
	}
}

func Benchmark_isPalindromeVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindromeVer2("A man, a plan, a canal: Panama")
	}
}

func Benchmark_isPalindromeVer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindromeVer3("A man, a plan, a canal: Panama")
	}
}
