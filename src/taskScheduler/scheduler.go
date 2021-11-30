package scheduler

import (
	"fmt"
)

type Priority int

const (
	PRIORI_UNKNOWN Priority = 0
	PRIORI_LOW = 1
	PRIORI_LOW_MEDIUM = 2
	PRIORI_MEDIUM = 3
	PRIORI_HIGH_MEDIUM = 4
	PRIORI_HIGH = 5
)

type TimeSlot struct {
	hour int
	minute int
	duration int
	free bool
}

type Tasklet struct {
	title string
	description string
	ts *TimeSlot
	chunks int
	priority Priority
}

func NewTimeSlot(hour, min, duration int) *TimeSlot {
	return &TimeSlot {
		hour: hour,
		minute: min,
		duration: duration,
		free: true,
	}
}

func NewTasklet(title string, desc string, hour int, min int, duration int, priori int, chunks int)  *Tasklet {
	return &Tasklet{
		title: title,
		description: desc,
		ts: NewTimeSlot(hour, min, duration),
		chunks: chunks,
		priority: Priority(priori),
	}
}

func (ts *TimeSlot) status() string {
	if ts.free {
		return "Free"
	}

	return "Blocked"
}

func (ts *TimeSlot) String() string {
	return fmt.Sprintf("%v Hours: %v minutes, for the duration of: %v is %s", ts.hour, ts.minute, ts.duration, ts.status())
}

func (tl *Tasklet) String() string {
	return fmt.Sprintf("%s: \n\t %q\n\t Time: %s\n\t Priority: %v\n\t Chunks: %v\n",
		tl.title, tl.description, tl.ts, tl.priority, tl.chunks)
}

func (tl *Tasklet) blockTimeSlot() {
	tl.ts.free = false
}