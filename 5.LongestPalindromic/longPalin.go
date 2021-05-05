package longpalin

func longestPalindrome(s string) string {
	n := len(s)
	vec := make([][]uint8, n)
	for i := 0; i < n; i++ {
		vec[i] = make([]uint8, n)
	}

	for i := 0; i < n; i++ {
		vec[i][i] = 1
	}

	start := 0
	maxLen := 1

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			start = i
			vec[i][i+1] = 1
			maxLen = 2
		}
	}

	for k := 3; k <= n; k++ {
		for i := 0; i < n-k+1; i++ {
			j := i + k - 1 // for indexing

			if (vec[i+1][j-1] == 1) && (s[i] == s[j]) {
				vec[i][j] = 1

				if k > maxLen {
					start = i
					maxLen = k
				}
			}

		}

	}

	return s[start : start+maxLen]
}

func longestPalindrome2(s string) string {
	var res, pa string
	var start int

	for start < len(s) {
		pa = findPalindrome(s, start, start)
		if len(pa) > len(res) {
			res = pa
		}

		pa = findPalindrome(s, start, start+1)
		if len(pa) > len(res) {
			res = pa
		}

		start++
	}

	return res
}

func findPalindrome(s string, left, right int) string {
	if len(s) == 0 || right >= len(s) || s[left] != s[right] {
		return ""
	}

	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}

	return s[left+1 : right]
}
