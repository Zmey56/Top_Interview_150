package intToRoman

import "testing"

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "Test 1",
			input:    3,
			expected: "III",
		},
		{
			name:     "Test 2",
			input:    4,
			expected: "IV",
		},
		{
			name:     "Test 3",
			input:    9,
			expected: "IX",
		},
		{
			name:     "Test 4",
			input:    58,
			expected: "LVIII",
		},
		{
			name:     "Test 5",
			input:    1994,
			expected: "MCMXCIV",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := intToRomain(test.input)
			if actual != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, actual)
			}
		})
	}
}

func Benchmark_intToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intToRomain(1994)
	}
}

func Benchmark_intToRomanV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intToRomainVer2(1994)
	}
}
