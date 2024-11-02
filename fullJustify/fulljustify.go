package fullJustify

import (
	"strings"
	"sync"
)

// Given an array of strings words and a width maxWidth, format the text such that each line has exactly maxWidth characters
// and is fully (left and right) justified.
//
// You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when
// necessary so that each line has exactly maxWidth characters.
//
// Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line does not divide
// evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.
//
// For the last line of text, it should be left-justified, and no extra space is inserted between words.
//
// Note:
//
// A word is defined as a character sequence consisting of non-space characters only.
// Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
// The input array words contains at least one word.

func fullJustify(words []string, maxWidth int) []string {
	result := []string{}

	for i := 0; i < len(words); {
		line := []string{}
		lineLength := 0

		for i < len(words) && lineLength+len(words[i])+len(line) <= maxWidth {
			line = append(line, words[i])
			lineLength += len(words[i])
			i++
		}

		if i == len(words) {
			result = append(result, leftJustify(line, maxWidth))
		} else {
			result = append(result, justify(line, maxWidth))
		}
	}

	return result
}

func leftJustify(line []string, maxWidth int) string {
	result := ""
	for i, word := range line {
		if i > 0 {
			result += " "
		}
		result += word
	}

	for i := len(result); i < maxWidth; i++ {
		result += " "
	}

	return result
}

func justify(line []string, maxWidth int) string {
	if len(line) == 1 {
		return leftJustify(line, maxWidth)
	}

	totalLength := 0
	for _, word := range line {
		totalLength += len(word)
	}

	spaces := maxWidth - totalLength
	gaps := len(line) - 1
	spacesBetweenWords := spaces / gaps
	extraSpaces := spaces % gaps

	result := ""
	for i, word := range line {
		if i > 0 {
			for j := 0; j < spacesBetweenWords; j++ {
				result += " "
			}
			if i <= extraSpaces {
				result += " "
			}
		}
		result += word
	}

	return result
}

func fullJustifyVer2(words []string, maxWidth int) []string {
	var result []string
	var currentLine []string
	lineLength := 0

	for _, word := range words {
		if lineLength+len(word)+len(currentLine) > maxWidth {
			spaces := maxWidth - lineLength
			for i := 0; i < spaces; i++ {
				currentLine[i%(len(currentLine)-1)] += " "
			}
			result = append(result, strings.Join(currentLine, ""))
			currentLine = nil
			lineLength = 0
		}
		currentLine = append(currentLine, word)
		lineLength += len(word)
	}

	result = append(result, strings.Join(currentLine, " ")+strings.Repeat(" ", maxWidth-lineLength-(len(currentLine)-1)))

	return result
}

func fullJustifyVer3(words []string, maxWidth int) []string {

	resultChan := make(chan string)
	var wg sync.WaitGroup
	var result []string

	go func() {
		for line := range resultChan {
			result = append(result, line)
		}
	}()

	processLine := func(lineWords []string, isLastLine bool) {
		defer wg.Done()
		line := ""
		if isLastLine || len(lineWords) == 1 {

			line = strings.Join(lineWords, " ") + strings.Repeat(" ", maxWidth-len(strings.Join(lineWords, " ")))
		} else {

			lineLength := 0
			for _, word := range lineWords {
				lineLength += len(word)
			}
			spaces := maxWidth - lineLength
			spaceSlots := len(lineWords) - 1
			for i := 0; i < spaceSlots; i++ {
				lineWords[i] += strings.Repeat(" ", spaces/spaceSlots)
				if i < spaces%spaceSlots {
					lineWords[i] += " "
				}
			}
			line = strings.Join(lineWords, "")
		}

		resultChan <- line
	}

	var currentLine []string
	lineLength := 0
	for _, word := range words {
		if lineLength+len(word)+len(currentLine) > maxWidth {

			wg.Add(1)
			go processLine(currentLine, false)
			currentLine = nil
			lineLength = 0
		}
		currentLine = append(currentLine, word)
		lineLength += len(word)
	}

	wg.Add(1)
	go processLine(currentLine, true)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	wg.Wait()
	return result
}
