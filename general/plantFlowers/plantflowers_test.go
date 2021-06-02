package plantflowers

import (
	"fmt"
	"testing"
)

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func TestPlantFlowers(t *testing.T) {
	in := []int {28, 23, 12, 37, 25, 9, 7, 22, 16}
	maxArea := []int{}

	for i := 1; i < len(in) - 1; i += 2 {
		maxArea = append(maxArea, max(max(in[i-1], in[i]), max(in[i], in[i+1])))
	}

	fmt.Println(maxArea)
}

func TestHeapPlantFlower(t *testing.T) {
	in := [][]int {
		{28, 23, 12, 37, 25, 9, 7, 22, 16},
		{28, 23, 12, 37, 9, 7, 22, 16},
		{28, 23, 12, 37, 25, 94, 9, 7, 22, 16},
		{14, 23, 12, 37, 25, 9, 7, 2, 6, 22, 16},
	}

	for _, v := range(in) {
		heapPlantFlower(v)
	}
}

func heapPlantFlower(in []int) {	
	fmt.Printf("Input Array: %+v\n", in)
	out := plantFlowers(in)
	fmt.Printf("Value Obtained: %+v\n", out)
	for _, v := range(out) {
		fmt.Printf("%v, \t", in[v])
	}

	fmt.Println()
}