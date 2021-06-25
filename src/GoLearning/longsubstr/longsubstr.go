package main

func lengthOfLongestSubstring(s string) int {
	var bitMap = make(map[byte]int, 26)
	var ans int
	var i int
	for j := 0; j < len(s); j++ {
		if bitMap[s[j]] > i {
			i = bitMap[s[j]]
		}

		if ans < j-i+1 {
			ans = j - i + 1
		}

		bitMap[s[j]] = j + 1

	}

	return ans
}
