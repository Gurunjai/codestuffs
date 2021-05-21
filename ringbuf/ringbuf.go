package ringbuf

import (
	"fmt"
	"runtime/debug"
)

// RingBuf - to provide the RingBuffer type define
type RingBuf struct {
	val []int
	size int
	free int
	frontIdx int
	rearIdx int
}

// NewRingBuffer - gets the new ring buffer object for the provided capactiy
func NewRingBuffer(cap int) *RingBuf {
	return &RingBuf {
		val : make([]int, cap),
		size : cap,
		free : cap,
		frontIdx : -1,
		rearIdx : -1,
	}
}

// IsEmpty - Check and see whether the RingBuffer is Empty
func (rb *RingBuf) IsEmpty() bool {
	return (rb.free == rb.size)
}

// IsFull - Check and see whether the RingBuffer reached its max capactiy
func (rb *RingBuf) IsFull() bool {
	return (rb.free == 0)
}

// Enqueue - Inserts the item into the ring buffer
func (rb *RingBuf) Enqueue(elem int) bool {
	if rb.IsFull() { return false }

	rb.rearIdx = ((rb.rearIdx + 1) % rb.size)

	rb.val[rb.rearIdx] = elem
	rb.free--

	return true
}

// Dequeue - Removes/extracts the item from the front of the RingBuffer, It also responds a boolean flag 
// to indicate whether the dequeue is success or failure
func (rb *RingBuf) Dequeue() (int, bool) {
	if rb.IsEmpty() { return -1, false }

	rb.frontIdx = ((rb.frontIdx + 1) % rb.size)
	rb.free++

	val := rb.val[rb.frontIdx]
	rb.val[rb.frontIdx] = -1

	return val, true
}

// Front - Get the Item in the front of the queue without removing it
func (rb *RingBuf) Front() int {
	if rb.IsEmpty() { return -1 }

	return rb.val[rb.frontIdx + 1]
}

// Rear - Get the item from the rear/latest entered into the queue
func (rb *RingBuf) Rear() int {
	if rb.IsEmpty() { return -1 }

	return rb.val[rb.rearIdx]
}

// Dump - Dump the RingBuffer elements
func (rb *RingBuf) Dump() {
	fmt.Printf("Values of the RingBuffer are: %+v\n", rb)
}

func init() {
	debug.SetTraceback("crash")

}