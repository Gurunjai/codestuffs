package numbits

import "testing"

func TestNumBits(t *testing.T) {
	in := map[uint32]int {
		0b00000000000000000000000000001011 : 3,
		0b00000000000000000000000010000000 : 1,
		0b00000000000000000000000010111000 : 4,
		0b00000000111111111111111111110000 : 20,
		0b11111111111111111111111111111101 : 31,
	}

	for s, want := range in {
		got := hammingWeight(s)

		if got != want {
			t.Fatalf("Failed for '%b':\n\t\t Got: %v \t Want: %v\n", s, got, want)
		}
	}
}