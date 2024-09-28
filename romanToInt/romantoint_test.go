package romanToInt

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Test 1",
			input:    "III",
			expected: 3,
		},
		{
			name:     "Test 2",
			input:    "IV",
			expected: 4,
		},
		{
			name:     "Test 3",
			input:    "IX",
			expected: 9,
		},
		{
			name:     "Test 4",
			input:    "LVIII",
			expected: 58,
		},
		{
			name:     "Test 5",
			input:    "MCMXCIV",
			expected: 1994,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := romanToInt(test.input)
			if actual != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, actual)
			}
		})
	}
}

func Benchmark_romanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		romanToInt("MCMXCIV")
	}
}

func Benchmark_romanToIntV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		romanToIntV2("MCMXCIV")
	}
}
