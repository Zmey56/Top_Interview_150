package strStr

import "strings"

// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack,
// or -1 if needle is not part of haystack.

func strStr(haystack string, needle string) int {
	if strings.ContainsAny(haystack, needle) {
		for i := 0; i < len(haystack); i++ {
			if haystack[i] == needle[0] {
				if strings.HasPrefix(haystack[i:], needle) {
					return i
				}
			}
		}
	}
	return -1
}

func strStrVer2(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func strStrVer3(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	lps := computeLPSArray(needle)
	i, j := 0, 0

	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
		}

		if j == len(needle) {
			return i - j
		} else if i < len(haystack) && haystack[i] != needle[j] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return -1
}

func computeLPSArray(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0
	i := 1

	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}
