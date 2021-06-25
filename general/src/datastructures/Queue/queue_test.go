package queue

import "testing"

func TestQueueBuffer( t *testing.T ) {
	buf := NewQueueBuffer(5)

	if ! buf.IsEmpty() {
		t.Fatalf("Queue Buffer is not empty!!!!\n")
	}

	in := []int { 10, 20, 30, 40, 50, 60, 70 }
	
	for _, n := range(in) {
		buf.Enqueue(n)
	}

	if ! buf.IsFull() {
		t.Fatalf("Queue Buffer is not full!!!!!\n")
	}

	if got, ok := buf.Rear(); ok && got != 50 {
		t.Errorf("Rear of the queue is not 50\n")
	}

	if got, ok := buf.Front(); ok && got != 10 {
		t.Errorf("Front of the queue is not 10")
	}

	buf.Dequeue()
	buf.Dequeue()

	buf.Enqueue(90)

	if got, ok := buf.Rear(); ok && got != 90 {
		t.Errorf("Rear of the queue is not 90\n")
	}

	buf.Dump()
}