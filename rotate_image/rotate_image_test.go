package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMinimalMatrix(t *testing.T) {
	matrix := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}
	expectedMatrix := [][]int{
		[]int{3, 1},
		[]int{4, 2},
	}

	rotate(matrix)

	assert.Equal(t, expectedMatrix, matrix)
}

func TestSwipeRight(t *testing.T) {
	matrix := [][]int{
		[]int{1, 0},
		[]int{0, 0},
	}
	expectedMatrix := [][]int{
		[]int{0, 1},
		[]int{0, 0},
	}
	swipe(0, 0, matrix)

	assert.Equal(t, expectedMatrix, matrix)
}

func TestSwipeDown(t *testing.T) {
	matrix := [][]int{
		[]int{0, 1},
		[]int{0, 0},
	}
	expectedMatrix := [][]int{
		[]int{0, 0},
		[]int{0, 1},
	}
	swipe(0,1, matrix)

	assert.Equal(t, expectedMatrix, matrix)
}

func TestSwipeLeft(t *testing.T) {
	matrix := [][]int{
		[]int{0, 0},
		[]int{0, 1},
	}
	expectedMatrix := [][]int{
		[]int{0, 0},
		[]int{1, 0},
	}
	swipe(1, 1, matrix)

	assert.Equal(t, expectedMatrix, matrix)
}

func TestSwipeUp(t *testing.T) {
	matrix := [][]int{
		[]int{0, 0},
		[]int{1, 0},
	}
	expectedMatrix := [][]int{
		[]int{1, 0},
		[]int{0, 0},
	}
	swipe(1, 0, matrix)

	assert.Equal(t, expectedMatrix, matrix)
}
