package convert

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display
// this pattern in a fixed font for better legibility)
//
// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"
//
// Write the code that will take a string and make this conversion given a number of rows:
//
// string convert(string s, int numRows);

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([][]byte, numRows)
	for i := range rows {
		rows[i] = make([]byte, 0)
	}

	down := false
	row := 0
	for i := 0; i < len(s); i++ {
		rows[row] = append(rows[row], s[i])
		if row == 0 || row == numRows-1 {
			down = !down
		}
		if down {
			row++
		} else {
			row--
		}
	}

	var result []byte
	for i := range rows {
		result = append(result, rows[i]...)
	}

	return string(result)
}

func convertVer2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	n := len(s)
	cycleLen := 2*numRows - 2
	result := make([]byte, 0, n)

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycleLen {
			result = append(result, s[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < n {
				result = append(result, s[j+cycleLen-i])
			}
		}
	}

	return string(result)
}

func convertVer3(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]chan byte, numRows)
	for i := range rows {
		rows[i] = make(chan byte, len(s))
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		down := false
		row := 0
		for i := 0; i < len(s); i++ {
			rows[row] <- s[i]
			if row == 0 || row == numRows-1 {
				down = !down
			}
			if down {
				row++
			} else {
				row--
			}
		}
		for _, ch := range rows {
			close(ch)
		}
	}()

	var result []byte
	for _, ch := range rows {
		for c := range ch {
			result = append(result, c)
		}
	}

	<-done
	return string(result)
}
