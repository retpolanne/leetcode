package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHappyPathWithObstacle(t *testing.T) {
	grid := [][]int{
		[]int{0, 0, 0},
		[]int{1, 1, 0},
		[]int{0, 0, 0},
		[]int{0, 1, 1},
		[]int{0, 0, 0},
	}
	assert.Equal(t, 10, shortestPath(grid, 0))
}

func TestHappyPathBreakObstacle(t *testing.T) {
	grid := [][]int{
		[]int{0, 0, 0},
		[]int{1, 1, 0},
		[]int{0, 0, 0},
		[]int{0, 1, 1},
		[]int{0, 0, 0},
	}
	assert.Equal(t, 6, shortestPath(grid, 1))
}

func TestDeadEnd(t *testing.T) {
	grid := [][]int{
		[]int{0, 1, 1},
		[]int{1, 1, 1},
		[]int{1, 0, 0},
	}
	assert.Equal(t, -1, shortestPath(grid, 1))
}
