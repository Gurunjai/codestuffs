package fibdp

import "testing"

func initialize() map[int]uint64 {
	return map[int]uint64 {
		1: 1,
		2: 1,
		4: 3,
		8: 21,
		16: 987,
		32: 2178309,
		33: 3524578,
		40: 102334155,
		56: 225851433717,
		84: 160500643816367088,
		96: 14787220707439219840,
		121: 790376311979428689,
		256: 17050816337285819451,
		512: 7623193360756408773,
		1024: 17158967189505497659,
		4196: 712541349587831693,
	}
}

func TestFibSequence(t *testing.T) {
	in := map[int]int {
		1: 1,
		2: 1,
		4: 3,
		8: 21,
		16: 987,
		32: 2178309,
		40: 102334155,
		// Unable to do this an kept on waiting
		// 50: 0,
	}

	for n, want := range(in) {
		got := fib(n)

		if got != want {
			t.Fatalf("failed for N: %v\n \t\t Got: \t%v, \t Want: \t%v\n", n, got, want)
		}
	}
}

func TestFibonacci(t *testing.T) {
	in := initialize()

	for n, want := range(in) {
		got := Fibonacci(n)

		if got != want {
			t.Fatalf("Failed for DP #N: %v\n \t\t Got: \t%v, \t Want: \t%v\n", n, got, want)
		}
	}

	// Cache.dump()
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

func BenchmarkFibonacci(b *testing.B) {
	in := initialize()

	for n, want := range(in) {
		got := Fibonacci(n)

		if got != want {
			b.Fatalf("Failed for DP for N: %v\n \t\t Got: \t%v, \t Want: \t%v\n", n, got, want)
		}
	}
}

func BenchmarkFibNormal(b *testing.B) {
	in := map[int]int {
		1: 1,
		2: 1,
		4: 3,
		8: 21,
		16: 987,
		32: 2178309,
		40: 102334155,
		// Unable to do this an kept on waiting
		// 50: 0,
	}
	
	for n, want := range(in) {
		got := fib(n)
	
		if got != want {
			b.Fatalf("Failed for N: %v\n \t\t Got: \t%v, \t Want: \t%v\n", n, got, want)
		}
	}
}

func TestFibCache(t *testing.T) {
	in := map[int]uint64 {
		8: TableSize,
		10: TableSize >> 1,
		25: TableSize << 1,
		16: TableSize >> 2,
		4: TableSize << 2,
		2: TableSize >> 3,
		9: TableSize << 3,
		69: TableSize >> 4,
		36: TableSize << 4,
		12: TableSize >> 5,
		29: TableSize << 5,
		198: 0,
	}

	// var head FibVal
	var root *FibVal = nil
	for k, v := range(in) {
		root = Insert(root, k, v)
	}

	for k, want := range(in) {
		node := Find(root, k)
		if node == nil {
			t.Fatalf("Unexpected node being nil!!!!")
		}

		if want != node.val {
			t.Fatalf("Failed for key: %v\n, \t\t Got: \t%v, \t Want: \t%v", k, node.val, want)
		}
	}
}

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