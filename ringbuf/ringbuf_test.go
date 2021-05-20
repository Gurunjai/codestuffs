package ringbuf

import "testing"

func TestRingBufBase(t *testing.T) {
	rb := NewRingBuffer(10)

	if rb.isFull() == true {
		t.Errorf("The empty buffer can not be full\n")
	}

	in := []int { 10, 20, 30, 40, 50, 60, 70, 80, 99, 100 }
	for _, n := range(in) {
		rb.enqueue(n)
	}

	if ! rb.isFull() {
		t.Errorf("Not Reporting Buffer is Full\n")
	}

	if v := rb.front(); v != 10 {
		t.Errorf("Mismatched at front:\n \t\t Expected: %v, \t Got: %v\n", 10, v)
	}

	if v := rb.rear(); v != 100 {
		t.Errorf("Mismatched at rear:\n \t\t Expected: %v, \t Got: %v\n", 50, v)
	}

	for _, n := range(in) {
		if v, ok := rb.dequeue(); ok && v != n {
			t.Errorf("Not reporting in proper order:\n \t\t Want: %v, \t Got: %v\n", n, v)
		}
	}

	if ! rb.isEmpty() {
		t.Errorf("Not Reporting Buffer is Empty\n")
	}
}

func TestRingBufRandom(t *testing.T) {
	rb := NewRingBuffer(5)
	
	rb.enqueue(5)
	rb.enqueue(6)
	if v, ok := rb.dequeue(); ok {
		if v != 5 {
			t.Errorf("Mismatched:\n \t\t Value Required: %v, \t Value Got: %v\n", 5, v)
		}
	}

	if v := rb.front(); v != 6 {
		t.Errorf("Not Reporting a front Value:\n \t\t Expected: %v, \t Got: %v\n", 6, v)
	}

	rb.enqueue(10)
	rb.enqueue(1)
	rb.enqueue(20)

	if v := rb.rear(); v != 20 {
		t.Errorf("Not Reporting a rear Value:\n \t\t Expected: %v, \t Got: %v\n", 20, v)
	}

	rb.enqueue(25)
	// rb.dump()
	if ok := rb.enqueue(12); ok {
		t.Errorf("Got a Success Instead of Failure")
	}

	if v, ok := rb.dequeue(); ok {
		if v != 6 {
			t.Errorf("Mismatched:\n \t\t Value Required: %v, \t Value Got: %v\n", 6, v)
		}
	}
	
	if v := rb.front(); v != 10 {
		t.Errorf("Not Reporting a rear Value:\n \t\t Expected: %v, \t Got: %v\n", 10, v)
	}
	
	if v := rb.rear(); v != 25 {
		t.Errorf("Not Reporting a rear Value:\n \t\t Expected: %v, \t Got: %v\n", 25, v)
	}
	
	// rb.dump()
}