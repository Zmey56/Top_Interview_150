package romanToInt

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
//
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
//
// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer.

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	res := 0
	prevSymbol := ""

	for _, c := range s {
		currSymbol := string(c)
		if prevSymbol == "I" && (currSymbol == "V" || currSymbol == "X") {
			res -= 2
		} else if prevSymbol == "X" && (currSymbol == "L" || currSymbol == "C") {
			res -= 20
		} else if prevSymbol == "C" && (currSymbol == "D" || currSymbol == "M") {
			res -= 200
		}
		res += m[currSymbol]
		prevSymbol = currSymbol
	}

	return res
}

func romanToIntV2(s string) int {
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0
	n := len(s)

	for i := 0; i < n; i++ {
		value := m[rune(s[i])]
		if i < n-1 && value < m[rune(s[i+1])] {
			res -= value
		} else {
			res += value
		}
	}

	return res
}
