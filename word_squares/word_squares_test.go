package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWorldsSimplestWordSquare(t *testing.T) {
	/*
	 * A B
	 * B A
	 */
	wordSquare := []string{"AB", "BA"}
	assert.True(t, isWordSquare(wordSquare))
}

func TestNotASquareVertically(t *testing.T) {
	/*
	 * A B C
	 * B A U
	 */

	// This is not a square
	words := []string{"AB", "BA", "CU"}
	assert.False(t, isWordSquare(words))
}

func TestNotASquareHorizontally(t *testing.T) {
	/*
	 * A B C
	 * B A
	 */

	// This is not a square
	words := []string{"ABC", "BA"}
	assert.False(t, isWordSquare(words))
}

func TestNotAWordSquare(t *testing.T) {
	/*
	 * A B A
	 * B A X
	 * C A A
	 */

	// This is not a word square, although it is a square
	words := []string{"AB", "BA", "C"}
	assert.False(t, isWordSquare(words))
}

func TestWordSameVertically(t *testing.T) {
	/*
	 * B A L L
	 * A R E A
	 * L E A D
	 * L A D Y
	 */

	words := []string{"BALL", "AREA", "LEAD", "LADY"}
	assert.True(t, wordSameVertically(1, words))
}

func TestWordNotSameVertically(t *testing.T) {
	/*
	 * B I L L
	 * A R E A
	 * L E A D
	 * L A D Y
	 */

	words := []string{"BALL", "IREA", "LEAD", "LADY"}
	assert.False(t, wordSameVertically(1, words))
}
