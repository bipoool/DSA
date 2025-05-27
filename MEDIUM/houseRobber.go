package medium

// At each index you have 2 options
// Either rob the 2nd house
// Or rob what amount is possible if last house was not robbed
func rob1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	prev := nums[1]
	for i := 2; i < len(nums); i++ {
		temp := max(nums[i-2], nums[i-1]-prev) + nums[i]
		prev = nums[i]
		nums[i] = temp
	}
	return max(nums[len(nums)-1], nums[len(nums)-2])
}
