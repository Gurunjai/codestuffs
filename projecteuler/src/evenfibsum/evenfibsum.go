package evenfibsum

import "math"

var (
	// SqrtOf5 - For Binet's Formula
	// http://www.maths.surrey.ac.uk/hosted-sites/R.Knott/Fibonacci/fibFormula.html#section1
	SqrtOf5 = math.Sqrt(5)

	// Phi - For Binet's Formula
	Phi = (1.0 + SqrtOf5) / 2
)

func fibonacci(index int) uint64 {
	return uint64(math.Round(1.0 / SqrtOf5 * math.Pow(Phi, float64(index))))
}

func sumOfEvenFibAtIndex(index int) uint64 {
	index = index / 3 * 3

	return (fibonacci(index + 2) - 1) / 2
}

func sumOfEvenFibonacci(bound int) uint64 {
	return sumOfEvenFibAtIndex(getTHeLargestIndexSmallerThanBound(bound))
}

func getTHeLargestIndexSmallerThanBound(bound int) int {
	return int(math.Floor( math.Log( SqrtOf5 * float64( bound ) ) / math.Log( Phi ) ) )
}