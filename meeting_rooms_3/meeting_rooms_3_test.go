package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestScheduleMeetingsFirstSlot(t *testing.T) {
	// Book room 0 - [2, 4)
	// Book room 0 - [0, 2)
	// Book room 1 - [2, 4)
	roomMgmt := &RoomMgmt{}
	// Start room pool with 2 rooms
	roomMgmt.Init(2, 3)

	meeting1 := &Meeting{}
	meeting1.Init(2, 4)

	meeting2 := &Meeting{}
	meeting2.Init(0, 2)

	meeting3 := &Meeting{}
	meeting3.Init(2, 4)

	// Meetings won't have any delay
	meeting1Ret := roomMgmt.TryToSchedule(meeting1)
	assert.Equal(t, 0, meeting1Ret.DelayedHours)
	meeting2Ret := roomMgmt.TryToSchedule(meeting2)
	assert.Equal(t, 0, meeting2Ret.DelayedHours)
	meeting3Ret := roomMgmt.TryToSchedule(meeting3)
	assert.Equal(t, 0, meeting3Ret.DelayedHours)

	room0Queue := []*Meeting{
		meeting1, meeting2, nil,
	}
	room1Queue := []*Meeting{
		meeting3, nil, nil,
	}
	assert.Equal(t, room0Queue, roomMgmt.Rooms[0].MeetingQueue)
	assert.Equal(t, room1Queue, roomMgmt.Rooms[1].MeetingQueue)
}

func TestScheduleMeetingsToTheLowestRoom(t *testing.T) {
	// Book room 0 - [1, 2)
	// Book room 1 - [0, 2)
	// Book room 0 - [2, 4)
	roomMgmt := &RoomMgmt{}
	// Start room pool with 2 rooms
	roomMgmt.Init(2, 3)

	meeting1 := &Meeting{}
	meeting1.Init(1, 2)

	meeting2 := &Meeting{}
	meeting2.Init(1, 2)

	meeting3 := &Meeting{}
	meeting3.Init(2, 4)

	// Meetings won't have any delay
	meeting1Ret := roomMgmt.TryToSchedule(meeting1)
	assert.Equal(t, 0, meeting1Ret.DelayedHours)
	meeting2Ret := roomMgmt.TryToSchedule(meeting2)
	assert.Equal(t, 0, meeting2Ret.DelayedHours)
	meeting3Ret := roomMgmt.TryToSchedule(meeting3)
	assert.Equal(t, 0, meeting3Ret.DelayedHours)

	room0Queue := []*Meeting{
		meeting1, meeting3, nil,
	}
	room1Queue := []*Meeting{
		meeting2, nil, nil,
	}
	assert.Equal(t, room0Queue, roomMgmt.Rooms[0].MeetingQueue)
	assert.Equal(t, room1Queue, roomMgmt.Rooms[1].MeetingQueue)
}

func TestScheduleMeetingsWithDelay(t *testing.T) {
	// Book room 0 - [1, 2)
	// Book room 1 - [1, 2)
	// Book room 0 - [1, 4)
	roomMgmt := &RoomMgmt{}
	// Start room pool with 2 rooms
	roomMgmt.Init(2, 3)

	meeting1 := &Meeting{}
	meeting1.Init(1, 2)

	meeting2 := &Meeting{}
	meeting2.Init(1, 2)

	meeting3 := &Meeting{}
	meeting3.Init(1, 4)

	// Meetings will have delay
	meeting1Ret := roomMgmt.TryToSchedule(meeting1)
	assert.Equal(t, 0, meeting1Ret.DelayedHours)
	meeting2Ret := roomMgmt.TryToSchedule(meeting2)
	assert.Equal(t, 0, meeting2Ret.DelayedHours)
	meeting3Ret := roomMgmt.TryToSchedule(meeting3)
	assert.Equal(t, 1, meeting3Ret.DelayedHours)

	room0Queue := []*Meeting{
		meeting1, meeting3, nil,
	}
	room1Queue := []*Meeting{
		meeting2, nil, nil,
	}
	assert.Equal(t, room0Queue, roomMgmt.Rooms[0].MeetingQueue)
	assert.Equal(t, room1Queue, roomMgmt.Rooms[1].MeetingQueue)
}

func TestMostBooked(t *testing.T) {
	meetings := [][]int{
		[]int{0, 10},
		[]int{1, 5},
		[]int{2, 7},
		[]int{3, 4},
	}
	assert.Equal(t, 0, mostBooked(2, meetings))
}

func TestMostBooked2(t *testing.T) {
	meetings := [][]int{
		[]int{1, 20},
		[]int{2, 10},
		[]int{3, 5},
		[]int{4, 9},
		[]int{6, 8},
	}
	assert.Equal(t, 1, mostBooked(3, meetings))
}

func TestMostBooked3(t *testing.T) {
	meetings := [][]int{
		[]int{10, 11},
		[]int{2, 10},
		[]int{1, 17},
		[]int{9, 13},
		[]int{18, 20},
	}
	assert.Equal(t, 1, mostBooked(2, meetings))
}

func TestMostBooked4(t *testing.T) {
	meetings := [][]int{
		[]int{4,11},
		[]int{1,13},
		[]int{8,15},
		[]int{9,18},
		[]int{0,17},
	}
	assert.Equal(t, 1, mostBooked(2, meetings))
}

func TestMostBooked5(t *testing.T) {
	meetings := [][]int{
		[]int{18,19},
		[]int{3,12},
		[]int{17,19},
		[]int{2,13},
		[]int{7,10},
	}
	assert.Equal(t, 0, mostBooked(4, meetings))
}

func TestMostBooked6(t *testing.T) {
	meetings := [][]int{
		[]int{48,50},
		[]int{38,41},
		[]int{17,48},
		[]int{10,26},
		[]int{16,46},
	}
	assert.Equal(t, 1, mostBooked(2, meetings))
}
