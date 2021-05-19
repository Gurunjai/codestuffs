package heap

// IntHeap - Wrapper for the Int slices
type IntHeap []int

var minHeap bool

// Len - to provide the size of the heap
func (h IntHeap) Len() int { return len(h) }

// Less - Interface method to define the possibility of Min/Max Heap
func (h IntHeap) Less(i, j int) bool { 
	if minHeap {
		return h[i] < h[j]
	}

	return h[i] > h[j]
}

// Swap - To swap the elements within the slice array
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push - Interface for container/heap to add a new element to the heap
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop - Retrieve the top of the heap
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func heapProperty(min bool) {
	minHeap = min
}

