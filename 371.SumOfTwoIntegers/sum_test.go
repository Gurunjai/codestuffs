package sumoftwoint

import "testing"

func TestGetSum(t *testing.T) {
	const WantIdx = 2
	const XIdx = 0
	const YIdx = 1
	in := map[string][]int {
		"first" : []int {1, 2, 3},
		"second" : []int {2, 3, 5},
		"third" : []int {10, 45, 55},
		"four" : []int {120, 250, 370},
	}

	for s, v := range in {
		want := v[WantIdx]
		got := getSum(v[XIdx], v[YIdx])

		if got != want {
			t.Fatalf("Failed for '%s':\n\t\t Got: %v \t Want: %v\n", s, got, want)
		}
	}
}