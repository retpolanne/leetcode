package main

import "fmt"

func main() {
	_ = []string{
		"able",
		"ale",
		"apple",
		"bale",
		"kangaroo",
	}
}

func findSubsequences(setOfWords []string, givenString string) []string {
	return []string{
		"able",
		"ale",
		"apple",
	}
}

func isSubsequence(givenString string, word string) bool {
	match := false
	matches := 0
	for _, c := range word {
		i := 0
		fmt.Printf("letter %s\n", string(c))
		for !match && i < len(givenString) {
			if string(c) == string(givenString[i]) {
				match = true
				matches = matches + 1
				break
			} else {
				i = i + 1
			}
		}
		match = false
	}

	if (matches == len(word)) {
		return true
	}
	return false
}
