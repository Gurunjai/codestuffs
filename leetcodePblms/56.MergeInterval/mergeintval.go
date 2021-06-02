package mergeintval

type IntervalHeap [][]int
const start int = 0
const end int = 1

func (h IntervalHeap) Len() int {
	return len(h)
}

func (h IntervalHeap) Less(i, j int) bool {
	return h[i][start] < h[j][start]
}

func (h IntervalHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntervalHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *IntervalHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func overlap(a, b []int) bool {
	if a[start] > b[end] {
		return false
	}

	if b[start] > a[end] {
		return false
	}

	return true
}