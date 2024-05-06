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
