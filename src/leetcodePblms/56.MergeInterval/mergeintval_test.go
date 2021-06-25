package mergeintval

import "fmt"
import "testing"
import "container/heap"

func TestMergeInterval( t *testing.T ) {
	/* in := [][]int {
		[]int {1, 3},
		[]int {2, 6},
		[]int {8, 10},
		[]int {15, 18},
	} */

	in := [][]int {
		[]int {1, 4},
		[]int {5, 8},
	}

	h := &IntervalHeap{}
	heap.Init(h)

	for _, v := range(in) {
		heap.Push(h, v)
	}

	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}

}