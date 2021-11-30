package day1

func TwoSum(nums []int, target int) bool {
	inMap := make(map[int]bool)
	for _, v := range(nums) {
		k := target - v
		if _, ok := inMap[k]; ok {
			return true
		}
		inMap[v] = true
	}

	return false
}