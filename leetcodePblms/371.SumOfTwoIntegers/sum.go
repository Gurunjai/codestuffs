package sumoftwoint

import "math"

func getSum(a, b int) int {
	return int(math.Log(math.Exp(float64(a)) * math.Exp(float64(b))))
}