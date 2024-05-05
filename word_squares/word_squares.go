package main

import (
	"strings"
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Hello world")
}

func isWordSquare(arrayOfWords []string) bool {
	// Assume it's a word square default
	// we'll prove it's not!
	wordSquare := true
	if (len(arrayOfWords) > 0) {
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
	}
	return wordSquare
}

func wordSameVertically(index int, arrayOfWords []string) bool {
	var sb strings.Builder
	for _, word := range arrayOfWords {
		sb.WriteString(string(word[index]))
	}

	if sb.String() == arrayOfWords[index] {
		return true
	}
	return false
}

func separateWordsWithSameLength(arrayOfWords []string) map[int][]string {
	wordLengths := map[int][]string{}
	for _, word := range arrayOfWords {
		wordLengths[len(word)] = append(
			wordLengths[len(word)],
			word,
		)
	}
	return wordLengths
}

func generatePossibleWordSquares(arrayOfWords []string) [][]string {
	possibleSquares := [][]string{}
	wordsWithSameLength := separateWordsWithSameLength(arrayOfWords)
	for _, wordListSameLength := range wordsWithSameLength {
		wordSquare := []string{}
		for i, word := range wordListSameLength {
			fmt.Println(word)
			// Iterate through words with same length
			if i > 0 {
				// We don't care about the zeroth char
				// we also don't care for the first word
				// We start extracting the vertical from the i=1,
				// which should be a single letter
				verticalWord := extractVertical(wordSquare, i)
				// If we didn't find a vertical word, then we should just
				// go to the next word!
				if verticalWord != "" {
					// Found vertical word, let's find all words that
					// start with this vertical word
					possibleWords := findWordsThatStartsWith(verticalWord, wordListSameLength, wordSquare)
					if len(possibleWords) > 0 {
						// TODO
						// we can have something like LEAD and LEED for example
						wordSquare = append(wordSquare, possibleWords[0])
					}
				}
			}
		}

		if (isWordSquare(wordSquare) && len(wordSquare) > 0) {
			possibleSquares = append(possibleSquares, wordSquare)
		}
	}
	return possibleSquares
}

func extractVertical(arrayOfWords []string, index int) string {
	var sb strings.Builder
	for _, word := range arrayOfWords {
		sb.WriteString(string(word[index]))
	}
	return sb.String()
}

func findWordsThatStartsWith(start string, arrayOfWords []string, ignore []string) []string {
	wordList := []string{}
	for _, word := range arrayOfWords {
		if len(word) < len(start) {
			continue
		}
		wordBeginning := word[0 : len(start)]
		if wordBeginning == start && !slices.Contains(ignore, word) {
			wordList = append(wordList, word)
		}
	}
	return wordList
}
