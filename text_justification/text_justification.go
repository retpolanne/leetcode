package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world")
	words := []string{"Compartmentalization", "is", "a", "form", "of", "psychological", "defense", "mechanism"}
	fullJustify(words, 20)
}


func fullJustify(words []string, maxWidth int) []string {
	textSlice := [][]string{}
	i := 0
	j := 0
	wordsLen := 0
	lineSlice := []string{}

	for {
		if (wordsLen > 0 && wordsLen < maxWidth) {
			// Account for spaces before words
			// except for the 0th word in the line
			wordsLen++
		}
		if (wordsLen <= maxWidth) && (i < len(words)) {
			nextWordLen := len(words[i]) + wordsLen
			// Check if word at pointer fits line
			if (nextWordLen <= maxWidth) {
				// Word fits! Let's add it to the builder
				lineSlice = append(lineSlice, words[i])
				wordsLen = nextWordLen
				i++
			} else {
				// Word doesn't fit, increment j and keep i in the same position!
				textSlice = append(textSlice, lineSlice)
				lineSlice = []string{}
				j++
				wordsLen = 0
			}
		} else {
			textSlice = append(textSlice, lineSlice)
			break
		}
	}
	return []string{}
}

func leftJustify(words []string, maxWidth int) string {
	rightPadding := maxWidth - len(words)
	var sb strings.Builder
	for _, word := range words {
		rightPadding = rightPadding - len(word)
		sb.WriteString(word)
		sb.WriteString(" ")
	}
	for _ = range rightPadding {
		sb.WriteString(" ")
	}

	return sb.String()
}
