package ringbuf

import "testing"

func TestRingBufBase(t *testing.T) {
	rb := NewRingBuffer(10)

	if rb.IsFull() == true {
		t.Errorf("The empty buffer can not be full\n")
	}

	in := []int { 10, 20, 30, 40, 50, 60, 70, 80, 99, 100 }
	for _, n := range(in) {
		rb.Enqueue(n)
	}

	if ! rb.IsFull() {
		t.Errorf("Not Reporting Buffer is Full\n")
	}

	if v := rb.Front(); v != 10 {
		t.Errorf("Mismatched at front:\n \t\t Expected: %v, \t Got: %v\n", 10, v)
	}

	if v := rb.Rear(); v != 100 {
		t.Errorf("Mismatched at rear:\n \t\t Expected: %v, \t Got: %v\n", 50, v)
	}

	for _, n := range(in) {
		if v, ok := rb.Dequeue(); ok && v != n {
			t.Errorf("Not reporting in proper order:\n \t\t Want: %v, \t Got: %v\n", n, v)
		}
	}

	if ! rb.IsEmpty() {
		t.Errorf("Not Reporting Buffer is Empty\n")
	}
}

func TestRingBufRandom(t *testing.T) {
	rb := NewRingBuffer(5)
	
	rb.Enqueue(5)
	rb.Enqueue(6)
	if v, ok := rb.Dequeue(); ok {
		if v != 5 {
			t.Errorf("Mismatched:\n \t\t Value Required: %v, \t Value Got: %v\n", 5, v)
		}
	}

	if v := rb.Front(); v != 6 {
		t.Errorf("Not Reporting a front Value:\n \t\t Expected: %v, \t Got: %v\n", 6, v)
	}

	rb.Enqueue(10)
	rb.Enqueue(1)
	rb.Enqueue(20)

	if v := rb.Rear(); v != 20 {
		t.Errorf("Not Reporting a rear Value:\n \t\t Expected: %v, \t Got: %v\n", 20, v)
	}

	rb.Enqueue(25)
	// rb.Dump()
	if ok := rb.Enqueue(12); ok {
		t.Errorf("Got a Success Instead of Failure")
	}

	if v, ok := rb.Dequeue(); ok {
		if v != 6 {
			t.Errorf("Mismatched:\n \t\t Value Required: %v, \t Value Got: %v\n", 6, v)
		}
	}
	
	if v := rb.Front(); v != 10 {
		t.Errorf("Not Reporting a rear Value:\n \t\t Expected: %v, \t Got: %v\n", 10, v)
	}
	
	if v := rb.Rear(); v != 25 {
		t.Errorf("Not Reporting a rear Value:\n \t\t Expected: %v, \t Got: %v\n", 25, v)
	}
	
	// rb.Dump()
}