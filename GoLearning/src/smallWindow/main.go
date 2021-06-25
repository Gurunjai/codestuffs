package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func main() {
	in1 := []int{3, 7, 5, 6, 9}
	l, r := window(in1)

	fmt.Printf("Left = %v, Right = %v\n", l, r)

}

func window(in []int) (int, int) {
	if in == nil || 0 >= len(in) {
		return 0, 0
	}

	var left, right, maxSeen, minSeen int

	for i := 0; i < len(in); i++ {
		maxSeen = max(in[i], maxSeen)
		if in[i] < maxSeen {
			right = i
		}
	}

	for j := len(in) - 1; j >= 0; j-- {
		minSeen = min(in[j], minSeen)
		if in[j] > minSeen {
			left = j
		}
	}

	return left, right
}