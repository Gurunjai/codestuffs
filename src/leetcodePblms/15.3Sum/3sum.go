package threesum

import "strconv"
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[0:n - 1]
    return x
    
}

func sort(nums []int) (out []int) {
    h := &IntHeap{}
    heap.Init(h)
    
    for _, v := range(nums) {
        heap.Push(h, v)
    }
    
    for h.Len() > 0 {
        out = append(out, heap.Pop(h).(int))
    }
    
    return    
}

func getHashKey(a, b, c int) string {
	return strconv.Itoa(a) + ":" + strconv.Itoa(b) + ":" + strconv.Itoa(c)
}

func threeSum(nums []int) [][]int {
    nums = sort(nums)
	result := make(map[string]bool)	
    out := [][]int {}
    
    for i, a := range(nums) {
        l, r := i+1, len(nums) - 1
        lLoop: for l < r {
            totalSum := a + nums[l] + nums[r]
            if totalSum > 0 {
                r -= 1
            } else if totalSum < 0 {
                l += 1
            } else {
				if _, ok := result[getHashKey(a, nums[l], nums[r])]; !ok {
					result[getHashKey(a, nums[l], nums[r])] = true
					out = append(out, []int{a, nums[l], nums[r]})
				}
				break lLoop
            }
            
        }
    }

    return out   
}
