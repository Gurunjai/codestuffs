package primefactor

func getPrimeList(count int) []int {
	primeList := []int{2}
	var num int = 3
	var cnt int
	
	for cnt < count{
		divisible := false
		inLoop: 
		for _, n := range primeList {
			if num % n == 0 { 
				divisible = true
				break inLoop 
			}
		}

		if ! divisible {
			primeList = append(primeList, num)
		}

		cnt++
		num += 2
	}

	return primeList
}

func getPrimeFactor(num int) int {
	primeList := getPrimeList(6000)
	out := []int{}

	for _, prime := range(primeList) {
		if num <= 0 {
			break
		}

		if num % prime == 0 {
			out = append(out, prime)
			num /= prime
		}
	}

	return out[len(out) - 1]
}