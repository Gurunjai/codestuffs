package palinnum

func palinNumber(x int) bool {
	c := 0
	t := x
	for t > 0 {
		c = (c * 10) + (t % 10)
		t /= 10
	}

	return (x == c)
}
