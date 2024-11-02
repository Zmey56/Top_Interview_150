package fullJustify

import (
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		want     []string
	}{
		{
			name:     "Example 1",
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			want:     []string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			name:     "Example 2",
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			want:     []string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			name:     "Example 3",
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			want:     []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.words, tt.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_fullJustify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20)
	}
}

func Benchmark_fullJustifyVer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fullJustifyVer2([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20)
	}
}

func Benchmark_fullJustifyVer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fullJustifyVer3([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20)
	}
}
