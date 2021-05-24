package stack

import "fmt"

// StackBuf - Defined type for the Stack Buffer
type StackBuf struct {
	val []int
	size int
	top int
}

// NewStackBuffer - to allocate the stack for the provided size
func NewStackBuffer(capacity int) *StackBuf {
	return &StackBuf{
		val: make([]int, capacity),
		size: capacity,
		top: -1,
	}
}

// IsEmpty - Whether the stack is empty or not
func (sb *StackBuf) IsEmpty() bool {
	return (sb.top == -1)
}

// IsFull - Whether the stack is full or not
func (sb *StackBuf) IsFull() bool {
	return (sb.top == sb.size - 1)
}

// Push - Insert the element onto the stack
func (sb *StackBuf) Push(elem int) bool {
	if sb.IsFull() { return false }

	sb.top++
	sb.val[sb.top] = elem

	return true
}

// Pop - Removes the item from the top of the stack
func (sb *StackBuf) Pop() (int, bool) {
	if sb.IsEmpty() { return -1, false }

	val := sb.val[sb.top]
	sb.val[sb.top] = -1
	sb.top--

	return val, false
}

// Top - Inspect the element in the top of the stack
func (sb *StackBuf) Top() (int, bool) {
	if sb.IsEmpty() { return -1, false }

	return sb.val[sb.top], true
}

// Dump - Print the stack buffer
func (sb *StackBuf) Dump() {
	fmt.Printf("The Stack Buffer looks like: %+v\n", sb)
}