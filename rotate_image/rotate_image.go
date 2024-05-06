package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

func rotate(matrix [][]int){}

func swipe(iStart, jStart int, matrix [][]int) {
	horBoundary := len(matrix[iStart]) - 1
	vertBoundary := len(matrix) - 1
	value := matrix[iStart][jStart]
	// This i, j pair is not a right corner
	if (jStart < horBoundary) {
		matrix[iStart][jStart + 1] = value
	} else if (jStart == horBoundary) {
		// top right corner
		matrix[jStart][jStart] = value
	} else if (iStart == vertBoundary) {
		// bottom right corner
		matrix[iStart][jStart - 1] = value
	} else {
		// bottom left corner
		matrix[iStart - 1][jStart - 1] = value
	}
	matrix[iStart][jStart] =  0

}
