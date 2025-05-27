package easy

// Just start from the start
// nums[i] = nums[i-1] + nums[i-2]
// 2 steps are allowed
func climbStairs(n int) int {

	if n == 1 {
		return 1
	}
	nums := make([]int, n+1)
	nums[1] = 1
	nums[2] = 2

	for i := 3; i <= n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums[n]
}
