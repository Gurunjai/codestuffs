package arrrotate

import "fmt"

func arrayRotate(nums []int, k int) {	
	reverse(nums, 0, len(nums) - 1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums) - 1)

	fmt.Printf("Nums: %+v\n", nums)	
}

func reverse(nums []int, start, end int) {
	i := start
	j := end
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}