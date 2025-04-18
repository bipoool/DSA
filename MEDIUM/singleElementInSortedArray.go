package medium

// Every element at even index should be the first element of the duplicates
// Every element at odd index should be the second element of the duplicates
func singleNonDuplicate(nums []int) int {

	l := 0
	r := len(nums) - 1

	for l < r {
		mid := (l + r) / 2
		if (mid+1)%2 == 1 {
			if mid < len(nums)-1 && nums[mid] == nums[mid+1] {
				l = mid + 1
			} else if mid > 0 && nums[mid] == nums[mid-1] {
				r = mid - 1
			} else {
				return nums[mid]
			}
		} else {
			if mid > 0 && nums[mid] == nums[mid-1] {
				l = mid + 1
			} else if mid < len(nums)-1 && nums[mid] == nums[mid+1] {
				r = mid - 1
			} else {
				return nums[mid]
			}
		}
	}
	return nums[l]
}
