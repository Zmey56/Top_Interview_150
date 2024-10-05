package lengthOfLastWord

import (
	"strings"
	"unicode/utf8"
)

// Given a string s consisting of words and spaces, return the length of the last word in the string.
//
// A word is a maximal
// substring
// consisting of non-space characters only.

func lengthOfLastWord(s string) int {
	arrWords := strings.Split(s, " ")

	for i := len(arrWords) - 1; i >= 0; i-- {
		if len(arrWords[i]) > 0 {
			return utf8.RuneCountInString(arrWords[i])
		}
	}

	return 0
}

func lengthOfLastWordVer2(s string) int {

	s = strings.TrimSpace(s)

	words := strings.Split(s, " ")

	return utf8.RuneCountInString(words[len(words)-1])
}

func lengthOfLastWordVer3(s string) int {
	s = strings.TrimSpace(s)

	lastSpaceIndex := strings.LastIndex(s, " ")

	return len(s) - lastSpaceIndex - 1
}

func lengthOfLastWordVer4(s string) int {
	length := 0
	inWord := false

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			inWord = true
			length++
		} else if inWord {
			break
		}
	}

	return length
}

func lengthOfLastWordVer5(s string) int {
	length := 0
	i := len(s) - 1

	for i >= 0 && s[i] == ' ' {
		i--
	}

	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}

	return length
}
