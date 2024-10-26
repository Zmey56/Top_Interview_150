package convert

import "testing"

func Test_convert(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		numRows int
		want    string
	}{
		{
			name:    "Example 1",
			s:       "PAYPALISHIRING",
			numRows: 3,
			want:    "PAHNAPLSIIGYIR",
		},
		{
			name:    "Example 2",
			s:       "PAYPALISHIRING",
			numRows: 4,
			want:    "PINALSIGYAHRPI",
		},
		{
			name:    "Example 3",
			s:       "A",
			numRows: 1,
			want:    "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.s, tt.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_convert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convert("PAYPALISHIRING", 3)
	}
}

func Benchmark_convertVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convertVer2("PAYPALISHIRING", 3)
	}
}

func Benchmark_convertVer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convertVer3("PAYPALISHIRING", 3)
	}
}
