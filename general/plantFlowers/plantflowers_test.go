package plantflowers

import (
	"fmt"
	"testing"
)


func heapPlantFlower(in []int) {	
	// fmt.Printf("Input Array: %+v\n", in)
	out := plantFlowers(in)
	// fmt.Printf("Value Obtained: %+v\n", out)
	var sum int
	for _, v := range(out) {
		// fmt.Printf("%v, \t", in[v])
		sum += in[v]
	}
	// fmt.Println()

	// fmt.Printf("Sum of the soil quality: %v\n", sum)
}

// Entirely Invalid to do this
func maxValPlantFlowers(in []int) []int {
	maxArea := make(map[int]bool)

	for i := 1; i < len(in) - 1; i++ {
		var maxIdx int
		if ((in[i - 1] > in[i]) && (in[i - 1] > in[i + 1])) {
			maxIdx = i - 1
		} else if ((in[i] > in[i - 1]) && (in[i] > in[i + 1])) {
			maxIdx = i
		} else {
			maxIdx = i + 1
		}

		if maxIdx != 0 {
			in[maxIdx - 1] = 0
		}

		if maxIdx != len(in) - 1 {
			in[maxIdx + 1] = 0
		}
 		
		if _, ok := maxArea[maxIdx]; !ok {
			maxArea[maxIdx] = true
		}
	}

	var out []int
	var sum int
	for k := range(maxArea) {
		out = append(out, k)
		sum += in[k]
	}

	// fmt.Printf("Sum of the soil quality: %v\n", sum)
	return out
}

func defineInputArray() [][]int {
	return [][]int {
		{28, 23, 12, 37, 25, 9, 7, 22, 16},
		{28, 23, 12, 37, 9, 7, 22, 16},
		{28, 23, 12, 37, 25, 94, 9, 7, 22, 16},
		{14, 23, 12, 37, 25, 9, 7, 2, 6, 22, 16},
	}	
}

func TestArrayPlant(t *testing.T) {
	in := defineInputArray()
	
	for _, v := range(in) {
		fmt.Println(maxValPlantFlowers(v))
	}
}

func TestHeapPlant(t *testing.T) {
	in := defineInputArray()

	for _, v := range(in) {
		heapPlantFlower(v)
	}
}

func BenchmarkArrayPlant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := defineInputArray()

		for _, v := range(in) {
			maxValPlantFlowers(v)
		}
	}
}

func BenchmarkHeapPlant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := defineInputArray()

		for _, v := range(in) {
			plantFlowers(v)
		}
	}
}