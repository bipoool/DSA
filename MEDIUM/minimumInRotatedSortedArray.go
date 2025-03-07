package medium

func findMin(nums []int) int {

	l := 0
	r := len(nums) - 1
	var mid int

	if nums[l] < nums[r] {
		return nums[l]
	}

	for l <= r {

		mid = (r + l) / 2
		if mid-1 >= 0 && nums[mid] < nums[mid-1] {
			break
		}
		if nums[mid] < nums[0] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return nums[mid]
}
