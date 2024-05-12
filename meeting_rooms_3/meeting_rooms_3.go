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
	Id int
	MeetingCount int
}

func (room *Room) Init(meetingsLen int, id int) {
	room.MeetingQueue = make([]*Meeting, meetingsLen)
	room.NextMeetingTime = 0
	// Putting a big number here because we want to compare with a lower number
	room.FirstMeetingTime = 100
	room.NextSlot = 0
	room.Id = id
	room.MeetingCount = 0
}

func (room *Room) ScheduleMeeting(meeting *Meeting) {
	fmt.Printf("Scheduling meeting that starts at %d ends at %d on room %d\n", meeting.Start, meeting.End, room.Id)
	fmt.Printf("next meeting time %d room %d\n", room.NextMeetingTime, room.Id)
	// This assumes that RoomMgmt already calculated the necessary delays
	if (meeting.Start >= room.NextMeetingTime) ||
		(meeting.End <= room.FirstMeetingTime) {
		// No delays required
		// meeting starts at the next available time slot
		// OR
		// Meeting ends before the first slot
		room.MeetingQueue[room.NextSlot] = meeting
		if (meeting.Start < room.FirstMeetingTime) {
			room.FirstMeetingTime = meeting.Start
		}
		if (meeting.End > room.NextMeetingTime) {
			room.NextMeetingTime = meeting.End
		}
		room.NextSlot++
		room.MeetingCount++
	}
	// The edge case here would be that RoomMgmt didn't calculate the correct delays
}

// Calculate if there will be any delays and return the delay
// Use this to compare which room will have the less delays
func (room *Room) CalculateDelayHelper(meeting *Meeting) int {
	delay := 0
	if (meeting.Start < room.NextMeetingTime) {
		delay = room.NextMeetingTime - meeting.Start
	}
	return delay
}

type RoomMgmt struct {
	SoonerAvailableRoom *Room
	Rooms []*Room
}

func (roomMgmt *RoomMgmt) Init(n int, meetingsLen int) {
	roomMgmt.Rooms = make([]*Room, n)
	for i := range n {
		roomMgmt.Rooms[i] = &Room{}
		roomMgmt.Rooms[i].Init(meetingsLen, i)
	}
}

func (roomMgmt *RoomMgmt) TryToSchedule(meeting *Meeting) *Meeting {
	// Use a huge delay to get the smallest one
	// {smallest delay, room id}
	smallestDelay := []int{1000, 0}

	for i := range roomMgmt.Rooms{
		if meeting.Start < roomMgmt.Rooms[i].NextMeetingTime {
			// We may have a delay in this room
			// we'll save on the room with less delay
			delay := roomMgmt.Rooms[i].CalculateDelayHelper(meeting)
			if (delay < smallestDelay[0]) {
				smallestDelay[0] = delay
				smallestDelay[1] = i
			}
		}
		if meeting.Start == roomMgmt.Rooms[i].NextMeetingTime {
			// Schedule right on, as we have space here
			roomMgmt.Rooms[i].ScheduleMeeting(meeting)
			return meeting
		}

		if meeting.End <= roomMgmt.Rooms[i].FirstMeetingTime {
			// We have a meeting that ends before the first meeting in the room
			// if we see this, we can schedule right on
			roomMgmt.Rooms[i].ScheduleMeeting(meeting)
			return meeting
		}
	}

	meeting.Start = meeting.Start + smallestDelay[0]
	meeting.End = meeting.End + smallestDelay[0]
	meeting.DelayedHours = smallestDelay[0]

	// Schedule to the room with the smallest delay
	roomMgmt.Rooms[smallestDelay[1]].ScheduleMeeting(meeting)

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
	// (roomId, count)
	mostBookedRoom := []int{-1, 0}
	roomMgmt := &RoomMgmt{}
	roomMgmt.Init(n, len(meetings))
	for _, meeting := range meetings {
		meet := &Meeting{}
		meet.Init(meeting[0], meeting[1])
		roomMgmt.TryToSchedule(meet)
	}

	for i, room := range roomMgmt.Rooms {
		if room.MeetingCount > mostBookedRoom[1] {
			mostBookedRoom[0] = i
			mostBookedRoom[1] = room.MeetingCount
		}
	}
	return mostBookedRoom[0]
}
