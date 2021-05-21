package heap

import (
	"container/heap"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	const maxVal int = 197
	h := &IntHeap{2, 1, maxVal, 5, 95, 10, 67}
	heap.Init(h)
	heap.Push(h, 3)

	if maxVal != (*h)[0] {
		t.Fatalf("Invalid Max Heap, Peek Index is not %v\n", maxVal)
	}

	if v := heap.Pop(h); v != maxVal {
		t.Fatalf("Got a different max value in Top: %v\n", v)
	}

	if maxVal == (*h)[0] {
		t.Fatalf("Invalid Max Heap, Peek is Still maxVal %v\n", maxVal)
	}
}

func TestMinHeap(t *testing.T) {
	const minVal int = 6
	heapProperty(true) // To have the Min Heap Property
	h := &IntHeap{22, 19, minVal, 65, 95, 10, 67}
	
	heap.Init(h)
	heap.Push(h, 93)

	if minVal != (*h)[0] {
		t.Fatalf("Invalid Min Heap, Peek Index is not %v\n", minVal)
	}

	if v := heap.Pop(h); v != minVal {
		t.Fatalf("Got a different Min value in Top: %v\n", v)
	}

	if minVal == (*h)[0] {
		t.Fatalf("Invalid Min Heap, Peek is Still minVal %v\n", minVal)
	}
}