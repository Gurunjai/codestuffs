package main

import "fmt"

func main() {
	lis := []int{2, -1, 3, 4, -7, 2, 1, 3, -4, -3}
	// lis := []int{2, -1, 3, 4, -10, 2, 1, 6, 4, -3}
	largestSumSubArray(lis)
}

func largestSumSubArray(in []int) {
	if len(in) <= 0 {
		panic("zero length array!!!!!")
	}
	var out []int
	max := 0
	maxIdx := 0
	stIdx := 0

	out = append(out, in[0])
	
	for i := 1; i < len(in); i++ {
		val := out[i-1] + in[i]
		out = append(out, val)
		if val > max {
			max = val
			maxIdx = i
		}
	}

	for i := maxIdx; ; i-- {
		if i == 0 || out[i - 1] <= 0 {
			stIdx = i
			break
		}
	}

	fmt.Printf("Input Array: \t%+v\n", in)
	fmt.Printf("Output Array: \t%+v\n", out)
	fmt.Printf("Max: %v, Start Index: %v, Max Index: %v\n", 
		max, stIdx, maxIdx)
}