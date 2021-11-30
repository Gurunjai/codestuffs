package sumoftwoint

import "math"

func getSum(a, b int) int {
	return int(math.Log(math.Exp(float64(a)) * math.Exp(float64(b))))
}

func addNums(a, b int) int {
	for b != 0 {
		tmp := (a & b) << 1
		a ^= b
		b = tmp
	}

	return a
}