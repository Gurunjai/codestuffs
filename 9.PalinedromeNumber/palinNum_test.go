package palinnum

import "testing"

func TestPalinNumber(t *testing.T) {
	in := map[int]bool{
		121:  true,
		123:  false,
		-121: false,
		111:  true,
		10:   false,
		-101: false,
		101:  true,
	}

	for k, want := range in {
		got := palinNumber(k)

		if got != want {
			t.Fatalf("Failed for %v:\n \t\t Got: %v \t Want: %v", k, got, want)
		}
	}
}
