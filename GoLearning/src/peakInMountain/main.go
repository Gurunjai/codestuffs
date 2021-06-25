package main

import "fmt"

func findMountainPeak(nums[] int) int {
	fmt.Printf("\tNums: %+v\n\tLength: %v\n", nums, len(nums))
	/* const MinInt = -int32(^uint32(0) >> 1) - 1
	ret := MinInt
	peak := MinInt

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i - 1] && nums[i] > nums[i + 1] {
			if int32(nums[i]) > peak {
				peak = int32(nums[i])
				ret = i
			}
		}
	} */

	l, r := 0, len(nums) - 1
	for l <= r {
		m := l + (r - l) / 2

		if nums[m] > nums[m - 1] && nums[m] > nums[m + 1] {
			return m
		} else if nums[m - 1] < nums[m] && nums[m] < nums[m + 1] {
			l = m + 1
		} else {
			r = m
		}
	}

	return -1
}
/* 
func main() {
	fmt.Println(findMountainPeak([]int{1, 2, 0}))
	fmt.Println(findMountainPeak([]int{24, 79, 100, 99, 98, 67, 24, 90, 80}))
	fmt.Println(findMountainPeak([]int{24,69,100,99,79,78,67,36,26,19}))
} */