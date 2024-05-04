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
	for _, c := range word {
		i := 0
		for !match && i < len(word) {
			fmt.Printf("%s\n", string(c))
			if string(c) == string(givenString[i]) {
				fmt.Printf(
					"letter %s matches given %s",
					string(c),
					string(givenString[i]),
				)
				match = true
			} else {
				i = i + 1
			}
		}
		if !match {
			break
		}
	}
	return match
}
