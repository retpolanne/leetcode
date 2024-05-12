package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

type Room struct {
	MeetingQueue []*Meeting
	NextMeetingTime int
	FirstMeetingTime int
	NextSlot int
}

func (room *Room) Init(meetingsLen int) {
	room.MeetingQueue = make([]*Meeting, meetingsLen)
	room.NextMeetingTime = 0
	// Putting a big number here because we want to compare with a lower number
	room.FirstMeetingTime = 100
	room.NextSlot = 0
}

func (room *Room) ScheduleMeeting(meeting *Meeting) {
	// This assumes that RoomMgmt already calculated the necessary delays
	if (meeting.Start >= room.NextMeetingTime) ||
		(meeting.End < room.FirstMeetingTime) {
		// No delays required
		// meeting starts at the next available time slot
		// OR
		// Meeting ends before the first slot
		// TODO since we don't care so much about ordering,
		// this meeting can be included next in the queue for now
		room.MeetingQueue[room.NextSlot] = meeting
		if (meeting.Start < room.FirstMeetingTime) {
			room.FirstMeetingTime = meeting.Start
		}
		if (meeting.End > room.NextMeetingTime) {
			room.NextMeetingTime = meeting.End
		}
		room.NextSlot++
	}
	// The edge case here would be that RoomMgmt didn't calculate the correct delays
}

type RoomMgmt struct {
	SoonerAvailableRoom *Room
	Rooms []*Room
}

func (roomMgmt *RoomMgmt) Init(n int, meetingsLen int) {
	roomMgmt.Rooms = make([]*Room, n)
	for i := range n {
		roomMgmt.Rooms[i] = &Room{}
		roomMgmt.Rooms[i].Init(meetingsLen)
	}
}

func (roomMgmt *RoomMgmt) TryToSchedule(meeting *Meeting) *Meeting {
	fmt.Printf("Will try to schedule meeting start %d end %d\n", meeting.Start, meeting.End)
	return meeting
}

type Meeting struct {
	OriginalStart int
	Start int
	End int
	DelayedHours int
}

func (meeting *Meeting) Init(start, end int) {
	meeting.OriginalStart = start
	// half-closed interval
	meeting.Start = start
	meeting.End = end
	meeting.DelayedHours = 0
}

func mostBooked(n int, meetings [][]int) int {
	return 0
}
