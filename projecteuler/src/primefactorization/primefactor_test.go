package primefactor

import "testing"

func TestPrimeList( t * testing.T ) {
	want := []int {
		2, 3, 5, 7, 11, 13, 17, 19,
	}

	got := getPrimeList(10)

	if len(got) != len(want) {
		t.Fatalf("Mismatch in Length of the arrays got and want\n")
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("Mistmatch detected for Index: %v\n \t\t Got: %v, \t Want: %v\n", i, got, want)
		}
	}
}

func TestPrimeFactor( t *testing.T ) {
	in := map[int]int {
		44: 11,
		600851475143: 6857,
	}

	for num, want := range(in) {
		if got := getPrimeFactor(num); got != want {
			t.Fatalf("Failed to obtain prime factor for %v:\n \t\t Got: %v, \t Want: %v\n", num, got, want)
		}
	}
}