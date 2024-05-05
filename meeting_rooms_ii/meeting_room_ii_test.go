package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMinMeetingRoomsNoOverlap(t *testing.T) {
	intervals := [][]int{
		[]int{7, 10},
		[]int{2, 4},
	}
	assert.Equal(t, minMeetingRooms(intervals), 1)
}

func TestMinMeetingRoomsWithOverlap(t *testing.T) {
	intervals := [][]int{
		[]int{0, 30},
		[]int{5, 10},
		[]int{15, 20},
	}
	assert.Equal(t, minMeetingRooms(intervals), 2)
}
