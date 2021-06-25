package evenfibsum

import "testing"

func TestFibEvenSum(t *testing.T) {
	in := map[int]uint64 {
		4000000: 4613732,
		3000: 3382,
		2000000: 1089154,
		9000000: 4613732,
	}

	for bound, want := range(in) {
		if got := sumOfEvenFibonacci(bound); got != want {
			t.Fatalf("Mismatched for %v:\n \t\t Want: %v, \t Got: %v\n", bound, want, got)
		}
	}	
}


func BenchmarkFibonacciBinet(b *testing.B) {
	in := map[int]uint64 {
		1: 1,
		2: 1,
		4: 3,
		8: 21,
		16: 987,
		32: 2178309,
		33: 3524578,
		40: 102334155,
		56: 225851433717,
		78: 8944394323791463,
	}

	for n, want := range(in) {
		got := fibonacci(n)

		if got != want {
			b.Fatalf("Failed for Binet for N: %v\n \t\t Got: \t%v, \t Want: \t%v\n", n, got, want)
		}
	}
}