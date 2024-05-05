package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

type MidPoint struct {
	i int
	j int
}

func numIslands(grid [][]byte) int {
	foundIslands := 0
	return foundIslands
}

type VisitedIndex struct {
	I int
	J int
}

// Return the iterations and the visited indexes
// Return something like this
//
// iteration 0: []VisitedIndex{(0,0)}
// iteration 1: []VisitedIndex{(1,0), (0,1)}
// iteration 2: []VisitedIndex{(2,0), (0,2), (1,1)}
// ...
func islandTraverse(grid [][]byte) [][]VisitedIndex {
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
	 * x 0 0
	 *
	 */
	endHor := len(grid[0]) - 1
	endVer := len(grid) - 1
	index := 0
	iteration := 0
	// For diagonal traversal
	midpoint := 0
	iteratorIndexes := [][]VisitedIndex{}

	// iteration loop
	for {
		visitedIndexes := []VisitedIndex{}
		// Exit when the midpoint == endHor && endVer
		// TODO: test for non-square grids
		if (midpoint == endHor) && (midpoint == endVer) {
			break
		}

		// workload
		if (index == midpoint) {
			visitedIndexes = append(
				visitedIndexes,
				VisitedIndex{
					I: midpoint,
					J: midpoint,
				},
			)
		} else if (index > midpoint) {
			visitedIndexes = append(
				visitedIndexes,
				VisitedIndex{
					I: index,
					J: midpoint,
				},
			)
			visitedIndexes = append(
				visitedIndexes,
				VisitedIndex{
					I: midpoint,
					J: index,
				},
			)
			midpoint = index - 1
		}
		index++
		fmt.Printf(
			"Visited indexes for iteration %d: %v\n",
			iteration,
			visitedIndexes,
		)
		// Last thing: append visitedIndexes to iteratorIndexes
		iteratorIndexes = append(iteratorIndexes, visitedIndexes)
		iteration++
	}
	return iteratorIndexes
}
