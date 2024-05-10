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
	textSlice := []string{}
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
				textSlice = append(textSlice, justifyCenter(lineSlice, maxWidth))
				lineSlice = []string{}
				j++
				wordsLen = 0
			}
		} else {
			textSlice = append(textSlice, leftJustify(lineSlice, maxWidth))
			break
		}
	}
	return textSlice
}

func justifyCenter(words []string, maxWidth int) string {
	lineLen := 0
	letterCount := findHowManyLetters(words)
	idealSpaceLen := 1
	spaceCount := len(words) - 1
	spaceAdded := 0
	smallLastSpace := false

	if (len(words) > 1) {
		remainingSpaces := (maxWidth - letterCount)
		if remainingSpaces % 2 == 0 {
			idealSpaceLen = remainingSpaces / (spaceCount)
		} else {
			idealSpaceLen = (remainingSpaces + 1) / (spaceCount)
		}
	}

	if idealSpaceLen * spaceCount + letterCount > maxWidth {
		smallLastSpace = true
	}

	var sb strings.Builder
	for i, word := range words {
		sb.WriteString(word)
		lineLen = lineLen + len(word)
		// Don't add a space after the last word
		if (i < len(words) - 1) && (len(word) < maxWidth) {
			if spaceAdded == spaceCount - 1 && smallLastSpace {
				for _ = range idealSpaceLen - 1 {
					sb.WriteString(" ")
					lineLen++
				}
			} else {
				for _ = range idealSpaceLen {
					sb.WriteString(" ")
					lineLen++
				}
			}
			spaceAdded++
		} else if (len(words) == 1) {
			return leftJustify(words, maxWidth)
		}
	}
	return sb.String()
}

func findHowManyLetters(words []string) int {
	letterCount := 0
	for _, word := range words {
		letterCount = letterCount + len(word)
	}
	return letterCount
}

func leftJustify(words []string, maxWidth int) string {
	rightPadding := maxWidth - len(words)
	var sb strings.Builder
	for _, word := range words {
		rightPadding = rightPadding - len(word)
		sb.WriteString(word)
		if (len(word) < maxWidth) {
			sb.WriteString(" ")
		}
	}
	for _ = range rightPadding {
		sb.WriteString(" ")
	}

	return sb.String()
}
