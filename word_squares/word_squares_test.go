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

func TestGeneratePossibleWordSquares(t *testing.T) {
	words := []string{"AREA", "BALL", "DEAR", "LADY", "LEAD", "YARD"}
	output := [][]string{
		[]string{"BALL", "AREA", "LEAD", "LADY"},
		[]string{"LADY", "AREA", "DEAR", "YARD"},
	}
	assert.Equal(t, generatePossibleWordSquares(words), output)
}

func TestGeneratePossibleWordSquaresPt2(t *testing.T) {
	words := []string{"AREA", "BALL", "DEAR", "LADY", "LEAD", "YARD", "GATINHO"}
	output := [][]string{
		[]string{"BALL", "AREA", "LEAD", "LADY"},
		[]string{"LADY", "AREA", "DEAR", "YARD"},
	}
	assert.Equal(t, output, generatePossibleWordSquares(words))
}

func TestGeneratePossibleWordSquaresPt3(t *testing.T) {
	words := []string{"AREA", "BALL", "DEAR", "LADY", "LEAD", "LEED", "YARD", "AB", "BA"}
	output := [][]string{
		[]string{"BALL", "AREA", "LEAD", "LADY"},
		[]string{"LADY", "AREA", "DEAR", "YARD"},
		[]string{"AB", "BA"},
	}
	assert.Equal(t, output, generatePossibleWordSquares(words))
}

func TestSeparateWordsWithSameLength(t *testing.T) {
	words := []string{"AREA", "BALL", "DEAR", "LADY", "LEAD", "YARD", "AB", "BA"}
	output := map[int][]string{
		4: []string{"AREA", "BALL", "DEAR", "LADY", "LEAD", "YARD"},
		2: []string{"AB", "BA"},
	}
	assert.Equal(t, output, separateWordsWithSameLength(words))
}

func TestFindWordThatStartsEmpty(t *testing.T) {
	words := []string{"AREA", "BALL", "DEAR", "LADY", "LEAD", "LEED", "YARD", "BA"}
	assert.Equal(t, []string{}, findWordsThatStartsWith("A", words, []string{"AREA"}))
}

func TestFindWordThatStartsWith(t *testing.T) {
	words := []string{"AREA", "BALL", "DEAR", "LADY", "LEAD", "LEED", "YARD", "AB", "BA"}
	assert.Equal(t, []string{"LEAD", "LEED"}, findWordsThatStartsWith("LE", words, []string{}))
}

func TestFindWordThatStartsWithIgnore(t *testing.T) {
	words := []string{"AREA", "BALL", "DEAR", "LADY", "LEAD", "LEED", "YARD", "AB", "BA"}
	assert.Equal(t, []string{"LEAD"}, findWordsThatStartsWith("LE", words, []string{"LEED"}))
}

func TestExtractVertical(t *testing.T) {
	words := []string{"BALL", "AREA"}
	assert.Equal(t, "LE", extractVertical(words, 2))
}

func TestExtractVerticalSingleItem(t *testing.T) {
	words := []string{"BALL"}
	assert.Equal(t, "A", extractVertical(words, 1))
}
