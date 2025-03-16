package medium

func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	forw := make([]int, len(nums))
	back := make([]int, len(nums))

	forw[0] = nums[0]
	forw[1] = nums[1]

	back[0] = nums[1]
	back[1] = nums[2]

	for i := 2; i < len(nums)-1; i++ {
		if forw[i-2] < forw[i-1]-nums[i-1] {
			forw[i] = forw[i-1] - nums[i-1] + nums[i]
		} else {
			forw[i] = forw[i-2] + nums[i]
		}
	}

	nums = append(nums, nums[0])
	nums = nums[1:]

	for i := 2; i < len(nums); i++ {
		if back[i-2] < back[i-1]-nums[i-1] {
			back[i] = back[i-1] - nums[i-1] + nums[i]
		} else {
			back[i] = back[i-2] + nums[i]
		}
	}

	return max(forw[len(forw)-1], forw[len(forw)-2], back[len(back)-1], back[len(back)-1])

}
