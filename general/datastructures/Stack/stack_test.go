package stack

import "testing"

func TestStackBuffer( t *testing.T ) {
	buf := NewStackBuffer(5)

	if ! buf.IsEmpty() {
		t.Fatalf("Stack Buffer is not empty!!!!\n")
	}

	in := []int { 10, 20, 30, 40, 50, 60, 70 }
	
	for _, n := range(in) {
		buf.Push(n)
	}

	if ! buf.IsFull() {
		t.Fatalf("Stack Buffer is not full!!!!!\n")
	}

	if got, ok := buf.Top(); ok && got != 50 {
		t.Errorf("Top of the stack is not 50\n")
	}

	buf.Pop()
	buf.Pop()

	buf.Push(90)

	if got, ok := buf.Top(); ok && got != 90 {
		t.Errorf("Top of the stack is not 90\n")
	}

	buf.Dump()
}