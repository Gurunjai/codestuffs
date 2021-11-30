package trapLava

func trapLava(nums []int) int {
	l, r := 0, len(nums) - 1
	var ans, lMax, rMax int
	for l < r {
		if nums[l] < nums[r] {
			if nums[l] >= lMax {
				lMax = nums[l]
			} else {
				ans += (lMax - nums[l])
			}

			l++
		} else {
			if nums[r] >= rMax {
				rMax = nums[r]
			} else {
				ans += (rMax - nums[r])
			}

			r--
		}
	}

	return ans
}