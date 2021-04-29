package reverse

import "testing"

func TestReverseInteger(t *testing.T) {
	in := map[int]int {
		123: 321,
		-123: -321,
		120: 21,
		0: 0,
	}

	for s, v := range in {
		want := v
		got := ReverseInteger(s)
		if got != want {
			t.Fatalf("Failed\n\t\t Got: %v \t Want: %v\n", got, want)
		}
	}
}