package hamming

import "testing"

func TestHammingDistance(t *testing.T) {
	const WantIndex = 2
	const XIndex = 0
	const YIndex = 1
	in := map[string][]int {
		"first" : []int {1, 4, 2},
		"second" : []int {3, 1, 1},
	}

	for s, v := range in {
		want := v[WantIndex]
		got := Distance(v[XIndex], v[YIndex])
		if got != want {
			t.Fatalf("Failed for '%v':\n\t\t Got: %v \t Want: %v", s, got, want)
		}
	}
}