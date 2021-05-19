package arrrotate

import "testing"

func TestRotateArray(t *testing.T) {
	in := map[int][][]int {
		1 : [][]int { {1, 2, 3, 4, 5, 6, 7}, {3} },
		2 : [][]int { {3, 99, -1, -100}, {2} },
	}

	for _, val := range in {
		nums := val[0]
		k := val[1][0]
		arrayRotate(nums, k)
	}
}