package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMinimalElevationMap(t *testing.T) {
	elevationMap := []int{1, 0, 1}

	assert.Equal(t, 1, trap(elevationMap))
}

func TestMinimalElevationMap2(t *testing.T) {
	// Boundaries won't trap the water
	elevationMap := []int{0, 1}

	assert.Equal(t, 0, trap(elevationMap))
}

func TestNegativeMap(t *testing.T) {
	elevationMap := []int{1, 0, 1}
	negativeMapExp := []int{0, 1, 0}

	negativeMap, waterVolume := buildNegativeMap(elevationMap)

	assert.Equal(t, negativeMapExp, negativeMap)
	assert.Equal(t, 1, waterVolume)
}

func TestMediumElevationMap2(t *testing.T) {
	elevationMap := []int{2, 1, 0, 1, 2}

	assert.Equal(t, 4, trap(elevationMap))
}

func TestNegativeMap2(t *testing.T) {
	elevationMap := []int{2, 1, 0, 1, 2}
	negativeMap := []int{0, 1, 2, 1, 0}

	negativeMap, waterVolume := buildNegativeMap(elevationMap)

	assert.Equal(t, negativeMap, negativeMap)
	assert.Equal(t, 4, waterVolume)
}

func TestFindMaxElevation(t *testing.T) {
	elevationMap := []int{3, 2, 1, 0, 1, 2, 3}

	assert.Equal(t, 3, maxElevation(elevationMap))
}
