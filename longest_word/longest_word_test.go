package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSubsequences(t * testing.T) {
	setOfWords := []string{
		"able",
		"ale",
		"apple",
		"bale",
		"kangaroo",
	}
	expectedSubsequences := []string{
		"able",
		"ale",
		"apple",
	}
	givenString := "abppplee"
	assert.Equal(
		t, expectedSubsequences,
		findSubsequences(setOfWords, givenString),
	)
}

func TestNoSubsequenceHard(t *testing.T) {
	// No way kangaroo will be a subsequence of abppplee
	givenString := "abppplee"
	testString := "kangaroo"
	assert.False(t, isSubsequence(givenString, testString))
}

func TestNoSubsequenceSoft(t *testing.T) {
	// The word apfel will have some letters matching,
	// but not others, such as f. Also, l is out of order
	givenString := "abppplee"
	testString := "apfel"
	assert.False(t, isSubsequence(givenString, testString))
}

func TestSubsequence(t *testing.T) {
	// All words from able match
	givenString := "abppplee"
	testString := "able"
	assert.True(t, isSubsequence(givenString, testString))
}
