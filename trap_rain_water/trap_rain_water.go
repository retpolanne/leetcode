package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

func maxElevation(height []int) int {
	maxElevation := 0
	for i := range height {
		if height[i] > maxElevation {
			maxElevation = height[i]
		}
	}
	return maxElevation
}

// Build water map
func buildNegativeMap(height []int) ([]int, int) {
	waterVolume := 0
	negativeMap := make([]int, len(height))
	max := maxElevation(height)
	for i := range height {
		// Borders will let water flow
		if (i > 0 && i < len(height) - 1) {
			negativeMap[i] = max - height[i]
			waterVolume = waterVolume + negativeMap[i]
		}
	}
	return negativeMap, waterVolume
}

func trap(height []int) int {
	_, waterVolume := buildNegativeMap(height)
	return waterVolume
}
