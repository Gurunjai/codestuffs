package queue

import "fmt"

// buffer - Data structure to keep track of queue buffer
type buffer struct {
	val []int
	size int
	free int
	front int
	rear int
}

// NewQueueBuffer - respond with a new queue instance for the provided capacity
func NewQueueBuffer(size int) *buffer {
	return &buffer {
		val: make([]int, size),
		size: size,
		free: size,
		front: -1,
		rear: -1,
	}
}

// IsEmpty - Checks whether the queue is empty
func (qb *buffer) IsEmpty() bool {
	return (qb.size == qb.free)
}

// IsFull - Checks whether the queue is full
func (qb *buffer) IsFull() bool {
	return (qb.free == 0)
}

// Enqueue - Adds the element to the latest on the queue
func (qb *buffer) Enqueue(elem int) bool {
	if qb.IsFull() { return false }

	qb.free--
	
	qb.rear = ((qb.rear + 1) % qb.size)
	qb.val[qb.rear] = elem
	return true
}

// Dequeue - Removes the element from the front of the queue
func (qb *buffer) Dequeue() (int, bool) {
	if qb.IsEmpty() { return -1, false }

	qb.free++

	qb.front = ((qb.front + 1) % qb.size)
	v := qb.val[qb.front]
	qb.val[qb.front] = -1
	return v, true
}

// Front - Return the first element pushed into the queue
func (qb *buffer) Front() (int, bool) {
	if qb.IsEmpty() { return -1, false }

	return qb.val[qb.front + 1], true
}

// Rear - Returns the latest element pushed into the queue
func (qb *buffer) Rear() (int, bool) {
	if qb.IsEmpty() { return -1, false }

	return qb.val[qb.rear], true
}

// Dump - Dumps the internals of queue
func (qb *buffer) Dump() {
	fmt.Printf("Values are: %+v\n", qb)
}