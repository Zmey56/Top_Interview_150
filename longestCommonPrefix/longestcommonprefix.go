package longestCommonPrefix

import "sort"

// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".

func longestCommonPrefix(strs []string) string {
	prefix := ""

	if len(strs) == 0 {
		return prefix
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return prefix
			}
		}
		prefix += string(strs[0][i])
	}
	return prefix
}

func longestCommonPrefixVer2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, str := range strs[1:] {
		for len(str) < len(prefix) || str[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

func longestCommonPrefixVer3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs)
	first, last := strs[0], strs[len(strs)-1]
	i := 0

	for i < len(first) && i < len(last) && first[i] == last[i] {
		i++
	}

	return first[:i]
}

func longestCommonPrefixVer4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	low, high := 1, minLen
	for low <= high {
		mid := (low + high) / 2
		if isCommonPrefix(strs, mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return strs[0][:(low+high)/2]
}

func isCommonPrefix(strs []string, length int) bool {
	prefix := strs[0][:length]
	for _, str := range strs[1:] {
		if len(str) < length || str[:length] != prefix {
			return false
		}
	}
	return true
}
