package lengthOfLastWord

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Test 1",
			input:    "Hello World",
			expected: 5,
		},
		{
			name:     "Test 2",
			input:    "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			name:     "Test 3",
			input:    "luffy is still joyboy",
			expected: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := lengthOfLastWord(test.input)
			if actual != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, actual)
			}
		})
	}
}

func Benchmark_lengthOfLastWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLastWord("luffy is still joyboy")
	}
}

func Benchmark_lengthOfLastWordVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLastWordVer2("luffy is still joyboy")
	}
}

func Benchmark_lengthOfLastWordVer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLastWordVer3("luffy is still joyboy")
	}
}

func Benchmark_lengthOfLastWordVer4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLastWordVer4("luffy is still joyboy")
	}
}

func Benchmark_lengthOfLastWordVer5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLastWordVer5("luffy is still joyboy")
	}
}
