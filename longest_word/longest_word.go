package main

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
	subsequences := []string{}
	for _, word := range setOfWords {
		if isSubsequence(givenString, word) {
			subsequences = append(subsequences, word)
		}
	}
	return subsequences
}

func isSubsequence(givenString string, word string) bool {
	matches := 0
	i := 0
	for _, c := range word {
		readGivenString := false
		for !readGivenString && i < len(givenString) {
			if string(c) == string(givenString[i]) {
				readGivenString = true
				matches = matches + 1
				break
			} else {
				i = i + 1
			}
		}
	}

	if (matches == len(word)) {
		return true
	}
	return false
}
