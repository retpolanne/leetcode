package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFullJustify(t *testing.T) {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	example := []string{
		"This    is    an",
		"example  of text",
		"justification.  ",
	}
	assert.Equal(t, example, fullJustify(words, 16))
}

func TestJustifySmallString(t *testing.T) {
	words := []string{"justification."}
	example := []string{
		"justification.  ",
	}
	assert.Equal(t, example, fullJustify(words, 16))
}

func TestJustifyTextWithBigString(t *testing.T) {
	words := []string{"Compartmentalization", "is", "a", "form", "of", "psychological", "defense", "mechanism"}
	example := []string{
		"Compartmentalization",
		"is    a    form   of",
		"psychological       ",
		"defense mechanism   ",
	}
	assert.Equal(t, example, fullJustify(words, 20))
}

func TestLeftJustifyLastLine(t *testing.T) {
	exampleLine := "defense mechanism   "
	assert.Equal(t, exampleLine, leftJustify([]string{"defense", "mechanism"}, 20))
}

func TestJustifyCenter(t *testing.T) {
	exampleLine := "is    a    form   of"

	assert.Equal(t, exampleLine, leftJustify([]string{"is", "a", "form", "of"}, 20))
}
