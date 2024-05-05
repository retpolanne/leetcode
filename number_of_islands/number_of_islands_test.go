package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSmolIsland(t *testing.T) {
	// Test a simple, small island
	/*
	 * 0 0 0
	 * 0 1 0
	 * 0 0 0
	 */
	grid := [][]byte{
		[]byte{0, 0, 0},
		[]byte{0, 1, 0},
		[]byte{0, 0, 0},
	}
	assert.Equal(t, numIslands(grid), 1)
}

func TestBestCaseIsland(t *testing.T) {
	// Test a simple, best case scenario, island
	/*
	 * 1 0
	 * 0 0
	 */
	grid := [][]byte{
		[]byte{1, 0},
		[]byte{0, 0},
	}
	assert.Equal(t, numIslands(grid), 1)
}

func TestWorstCaseIsland(t *testing.T) {
	// Test a simple, worst case scenario, island
	/*
	 * 0 0
	 * 0 1
	 */
	grid := [][]byte{
		[]byte{0, 0},
		[]byte{0, 1},
	}
	assert.Equal(t, numIslands(grid), 1)
}

func TestAdjacentIsland(t *testing.T) {
	// Test a simple, adjacent, island
	/*
	 * 0 0
	 * 1 1
	 */
	grid := [][]byte{
		[]byte{0, 0},
		[]byte{1, 1},
	}
	assert.Equal(t, numIslands(grid), 1)
}

// TODO test a bigger grid
func TestIslandTraversal(t *testing.T) {
	// Test these iterations
	/*
	 * Island traversal should be like this
	 * iteration 0
	 *
	 * 0
	 *
	 * iteration 1
	 *
	 * x 0
	 * 0
	 *
	 * iteration 2
	 *
	 * x x 0
	 * x 1
	 * 0
	 *
	 * iteration 3
	 *
	 * x x x
	 * x x 0
	 * x 0
	 *
	 * iteration 4
	 *
	 * x x x
	 * x x x
	 * x x 0
	 *
	 */
	grid := [][]byte{
		[]byte{0, 0, 0},
		[]byte{0, 1, 0},
		[]byte{0, 0, 0},
	}
	visitedIterations := [][]VisitedIndex{
		// iteration 0
		[]VisitedIndex{
			VisitedIndex{
				I: 0,
				J: 0,
			},
		},
		// iteration 1
		[]VisitedIndex{
			VisitedIndex{
				I: 1,
				J: 0,
			},
			VisitedIndex{
				I: 0,
				J: 1,
			},
		},
		// iteration 2
		[]VisitedIndex{
			VisitedIndex{
				I: 2,
				J: 0,
			},
			VisitedIndex{
				I: 0,
				J: 2,
			},
			VisitedIndex{
				I: 1,
				J: 1,
			},
		},
		// iteration 3
		[]VisitedIndex{
			VisitedIndex{
				I: 2,
				J: 1,
			},
			VisitedIndex{
				I: 1,
				J: 2,
			},
			VisitedIndex{
				I: 2,
				J: 2,
			},
		},
	}
	assert.Equal(t, islandTraverse(grid), visitedIterations)
}
