package main

func twoSum(nums []int, target int) []int {
	listMap := make(map[int]int, len(nums))
	var out []int
	for idx, val := range nums {
		listMap[val] = idx
	}

	for i := 0; i < len(nums); i++ {
		lookup := target - nums[i]
		idx, ok := listMap[lookup]

		if idx == i {
			out = []int{i, idx}
			continue
		}

		if ok {
			out = []int{i, idx}
			break
		}

	}

	return out
}
