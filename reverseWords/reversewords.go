package reverseWords

import "strings"

// Given an input string s, reverse the order of the words.
//
// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
//
// Return a string of the words in reverse order concatenated by a single space.
//
// Note that s may contain leading or trailing spaces or multiple spaces between two words.
// The returned string should only have a single space separating the words. Do not include any extra spaces.

func reverseWords(s string) string {
	sTrim := strings.TrimSpace(s)
	sSplit := strings.Split(sTrim, " ")
	result := ""

	for i := len(sSplit) - 1; i >= 0; i-- {
		if sSplit[i] != "" {
			result += sSplit[i] + " "
		}
	}

	return strings.TrimSpace(result)
}

func reverseWordsVer2(s string) string {
	sTrim := strings.TrimSpace(s)
	sSplit := strings.Fields(sTrim)
	var result strings.Builder

	for i := len(sSplit) - 1; i >= 0; i-- {
		result.WriteString(sSplit[i])
		if i != 0 {
			result.WriteString(" ")
		}
	}

	return result.String()
}

func reverseWordsVer3(s string) string {
	sTrim := strings.TrimSpace(s)
	sSplit := strings.Fields(sTrim)
	result := make([]string, 0)

	for i := len(sSplit) - 1; i >= 0; i-- {
		result = append(result, sSplit[i])
	}

	return strings.Join(result, " ")
}

func reverseWordsVer4(s string) string {
	sTrim := strings.TrimSpace(s)
	sSplit := strings.Fields(sTrim)
	result := make([]string, len(sSplit))

	for i := len(sSplit) - 1; i >= 0; i-- {
		result[len(sSplit)-1-i] = sSplit[i]
	}

	return strings.Join(result, " ")
}

func reverseWordsVer5(s string) string {
	s = strings.TrimSpace(s)

	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
