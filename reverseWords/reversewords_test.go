package reverseWords

import "testing"

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "the sky is blue",
			want: "blue is sky the",
		},
		{
			name: "Example 2",
			s:    "  hello world  ",
			want: "world hello",
		},
		{
			name: "Example 3",
			s:    "a good   example",
			want: "example good a",
		},
		{
			name: "Example 4",
			s:    "  Bob    Loves  Alice   ",
			want: "Alice Loves Bob",
		},
		{
			name: "Example 5",
			s:    "Alice does not even like bob",
			want: "bob like even not does Alice",
		},
		{
			name: "Example 6",
			s:    "  Bob    Loves  Alice   ",
			want: "Alice Loves Bob",
		},
		{
			name: "Example 7",
			s:    "  Bob    Loves  Alice   ",
			want: "Alice Loves Bob",
		},
		{
			name: "Example 8",
			s:    "  Bob    Loves  Alice   ",
			want: "Alice Loves Bob",
		},
		{
			name: "Example 9",
			s:    "  Bob    Loves  Alice   ",
			want: "Alice Loves Bob",
		},
		{
			name: "Example 10",
			s:    "  Bob    Loves  Alice   ",
			want: "Alice Loves Bob",
		},
		{
			name: "Example 11",
			s:    "  Bob    Loves  Alice   ",
			want: "Alice Loves Bob",
		},
		{
			name: "Example 12",
			s:    "  Bob    Loves  Alice   ",
			want: "Alice Loves Bob",
		},
		{
			name: "Example 13",
			s:    "  Bob    Loves  Alice   ",
			want: "Alice Loves Bob",
		},
		{
			name: "Example 14",
			s:    "  Bob    Loves  Alice   ",
			want: "Alice Loves Bob",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_reverseWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWords("Alice does not even like bob")
	}
}

func Benchmark_reverseWordsVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWordsVer2("Alice does not even like bob")
	}
}

func Benchmark_reverseWordsVer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWordsVer3("Alice does not even like bob")
	}
}

func Benchmark_reverseWordsVer4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWordsVer4("Alice does not even like bob")
	}
}

func Benchmark_reverseWordsVer5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseWordsVer5("Alice does not even like bob")
	}
}
