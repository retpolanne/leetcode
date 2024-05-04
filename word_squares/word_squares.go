package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func isWordSquare(arrayOfWords []string) bool {
	// Assume it's a word square default
	// we'll prove it's not!
	wordSquare := true
	wordLen := len(arrayOfWords[0])
	for i, word := range arrayOfWords {
		if (wordLen != len(word)) {
			wordSquare = false
			break
		}

		if (wordLen != len(arrayOfWords)) {
			wordSquare = false
			break
		} else {
			// It's a square, now let's check the words
			wordSquare = wordSameVertically(i, arrayOfWords)
			if !wordSquare {
				break
			}
		}
	}
	return wordSquare
}

func wordSameVertically(index int, arrayOfWords []string) bool {
	var sb strings.Builder
	for _, word := range arrayOfWords {
		sb.WriteString(string(word[index]))
	}

	fmt.Printf("Is %s == %s\n", sb.String(), arrayOfWords[index])
	if sb.String() == arrayOfWords[index] {
		return true
	}
	return false
}
