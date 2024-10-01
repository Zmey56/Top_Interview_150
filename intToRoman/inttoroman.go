package intToRoman

// Seven different symbols represent Roman numerals with the following values:
//
// Symbol	Value
// I	1
// V	5
// X	10
// L	50
// C	100
// D	500
// M	1000
// Roman numerals are formed by appending the conversions of decimal place values from highest to lowest. Converting a decimal place value into a Roman numeral has the following rules:
//
// If the value does not start with 4 or 9, select the symbol of the maximal value that can be subtracted from the input, append that symbol to the result, subtract its value, and convert the remainder to a Roman numeral.
// If the value starts with 4 or 9 use the subtractive form representing one symbol subtracted from the following symbol, for example, 4 is 1 (I) less than 5 (V): IV and 9 is 1 (I) less than 10 (X): IX. Only the following subtractive forms are used: 4 (IV), 9 (IX), 40 (XL), 90 (XC), 400 (CD) and 900 (CM).
// Only powers of 10 (I, X, C, M) can be appended consecutively at most 3 times to represent multiples of 10. You cannot append 5 (V), 50 (L), or 500 (D) multiple times. If you need to append a symbol 4 times use the subtractive form.
// Given an integer, convert it to a Roman numeral.

func intToRomain(num int) string {
	m := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	res := ""
	for num > 0 {
		if num >= 1000 {
			res += m[1000]
			num -= 1000
		} else if num >= 900 {
			res += m[900]
			num -= 900
		} else if num >= 500 {
			res += m[500]
			num -= 500
		} else if num >= 400 {
			res += m[400]
			num -= 400
		} else if num >= 100 {
			res += m[100]
			num -= 100
		} else if num >= 90 {
			res += m[90]
			num -= 90
		} else if num >= 50 {
			res += m[50]
			num -= 50
		} else if num >= 40 {
			res += m[40]
			num -= 40
		} else if num >= 10 {
			res += m[10]
			num -= 10
		} else if num >= 9 {
			res += m[9]
			num -= 9
		} else if num >= 5 {
			res += m[5]
			num -= 5
		} else if num >= 4 {
			res += m[4]
			num -= 4
		} else {
			res += m[1]
			num--
		}
	}

	return res

}

func intToRomainVer2(num int) string {
	values := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	res := ""
	for _, v := range values {
		for num >= v.value {
			res += v.symbol
			num -= v.value
		}
	}

	return res
}
