package plantflowers

import (
	"container/heap"
)

// FlowerBed - Defined Type to hold the heap value and position in the array
type FlowerBed struct {
	val int
	pos int
}

// FlowerBedHeap - Underlying interface type for the flower bed type
type FlowerBedHeap []*FlowerBed
/* 
// IntArr - Abstract type to provide the Integer Array for Neighbor validation
type IntArr []int
 */
func newFlowerBed(v, p int) *FlowerBed {
	return &FlowerBed {
		val: v,
		pos: p,
	}
}

// Len - Provides the Length of the Heap
func (fh FlowerBedHeap) Len() int {
	return len(fh)
}

// Less - Provides the value at index i is less/more than value at index j
// Used Greater than to depict the MAX Heap rather than MIN Heap
func (fh FlowerBedHeap) Less(i, j int) bool {
	return fh[i].val > fh[j].val
}

// Swap - Swap the value in the provided index i, j position in the heap
func (fh FlowerBedHeap) Swap(i, j int) {
	fh[i], fh[j] = fh[j], fh[i]
}

// Push - Pushes the value into the Heap at appropriate level
func (fh *FlowerBedHeap) Push(x interface{}) {
	*fh = append(*fh, x.(*FlowerBed))
}

// Pop - Remove the value from the top of the heap
func (fh *FlowerBedHeap) Pop() interface{} {
	old := *fh
	n := len(old)
	x := old[n-1]
	*fh = old[:n-1]
	return x
}

// Peek - Inspect the top of the heap
func (fh *FlowerBedHeap) Peek() interface{} {
	old := *fh
	n := len(old)
	x := old[n-1]
	return x
}

/* // Push - Push or append the integer position value to the IntArr
func (ia *IntArr) Push(pos int) {
	*ia = append(*ia, pos)
}

// IsNeighbor - Used to check whether the provided index is a neighbor of the index value in IntArr
func (ia *IntArr) IsNeighbor(i int) bool {
	for _, v := range(*ia) {
		if i == v+1 || i == v-1 {
			return true
		}
	}

	return false
}

// ToIntArray - populate the IntArray to a type of int slice
func (ia *IntArr) ToIntArray() (out []int) {
	for _, v := range(*ia) {
		out = append(out, v)
	}
	return
}
 */
 
func plantFlowers(arr []int) []int {
	h := &FlowerBedHeap{}
	out := []int{}
	heap.Init(h)

	for i, val := range(arr) {
		heap.Push(h, newFlowerBed(val, i))
	}

	oLoop: for h.Len() > 0 {
		v := heap.Pop(h).(*FlowerBed)
		for _, p := range(out) {
			if v.pos == p + 1 || v.pos == p - 1{
				continue oLoop
			}
		}
		
		out = append(out, v.pos)
	}

	return out
}
