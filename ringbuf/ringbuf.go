package ringbuf

import "fmt"

type RingBuf struct {
	val []int
	size int
	free int
	frontIdx int
	rearIdx int
}

func NewRingBuffer(cap int) *RingBuf {
	return &RingBuf {
		val : make([]int, cap),
		size : cap,
		free : cap,
		frontIdx : -1,
		rearIdx : -1,
	}
}

func (rb *RingBuf) isEmpty() bool {
	return (rb.free == rb.size)
}

func (rb *RingBuf) isFull() bool {
	return (rb.free == 0)
}

func (rb *RingBuf) enqueue(elem int) bool {
	if rb.isFull() { return false }

	rb.rearIdx = ((rb.rearIdx + 1) % rb.size)

	rb.val[rb.rearIdx] = elem
	rb.free--

	return true
}

func (rb *RingBuf) dequeue() (int, bool) {
	if rb.isEmpty() { return -1, false }

	rb.frontIdx = ((rb.frontIdx + 1) % rb.size)
	rb.free++

	val := rb.val[rb.frontIdx]
	rb.val[rb.frontIdx] = -1

	return val, true
}

func (rb *RingBuf) front() int {
	if rb.isEmpty() { return -1 }

	return rb.val[rb.frontIdx + 1]
}

func (rb *RingBuf) rear() int {
	if rb.isEmpty() { return -1 }

	return rb.val[rb.rearIdx]
}

func (rb *RingBuf) dump() {
	fmt.Printf("Values of the RingBuffer are: %+v\n", rb)
}