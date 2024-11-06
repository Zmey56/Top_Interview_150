package isSubsequence

// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
//
// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of
// the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence
// of "abcde" while "aec" is not).

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	sliceS := []rune(s)
	sliceT := []rune(t)

	for i := 0; i < len(sliceT); i++ {
		if sliceS[0] == sliceT[i] {
			sliceS = sliceS[1:]
			if len(sliceS) == 0 {
				return true
			}
		}
	}

	return false
}

func isSubsequenceVer2(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	i, j := 0, 0
	for j < len(t) {
		if s[i] == t[j] {
			i++
			if i == len(s) {
				return true
			}
		}
		j++
	}

	return false
}

func isSubsequenceVer3(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	result := make(chan bool)
	defer close(result)

	go func() {
		i := 0
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				i++
				if i == len(s) {
					result <- true
					return
				}
			}
		}
		result <- false
	}()

	return <-result
}
