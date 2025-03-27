package medium

func findDuplicate(nums []int) int {

	l := nums[0]
	r := nums[0]

	for true {
		l = nums[l]
		r = nums[nums[r]]
		if l == r {
			break
		}
	}

	l = nums[0]
	for l != r {
		l = nums[l]
		r = nums[r]
	}

	return l
}
