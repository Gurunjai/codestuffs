package multiplesof3

func sumOfMultiplesof3And5(bound int) int {
	var sum int

	sum += (3 + 5 + 6 + 9)
	for i := 10; i < bound; i++ {
		if ( ( i % 3 == 0 ) || ( i % 5 == 0 ) ) {
			sum += i
		}
	}

	return sum
}