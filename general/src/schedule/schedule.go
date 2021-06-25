package schedule

import (
	// "fmt"
)

type SlotStatus int

const (
	Unknown SlotStatus = iota
	Used
	Partial
	Free
)

const (
	MaxFreeMinutes int = 60
	MinFreeMinutes int = 1
)

var (
	DefaultDurationMinutes int = 5
	MaxDurationMinutes int = MaxFreeMinutes	
)

type Task struct {
	startTime int
	duration int
	description string
	next *Task
	prev *Task
}

func NewTask(startTime, duration int, desc string) *Task {
	return &Task {
		startTime: startTime,
		duration: duration,
		description: desc,
		next: nil,
		prev: nil,
	}
}

type TaskList struct {
	head *Task
	tail *Task
}

type TimeSlot struct {
	startTime int
	status SlotStatus
	freeMins int
	allocated TaskList	
}