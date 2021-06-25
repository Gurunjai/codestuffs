package hamming

// Distance this function provides the minimum number of bits different in two integers
func Distance(x, y int) int {
	c := x ^ y
	cnt := 0

	for c > 0 {
		c &= c - 1
		cnt++
	}

	return cnt
}