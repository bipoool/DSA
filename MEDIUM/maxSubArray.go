package medium

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	runSum := nums[0]
	maxSum := math.MinInt
	for i := 1; i < len(nums); i++ {
		runSum = max(runSum+nums[i], nums[i])
		maxSum = max(maxSum, runSum)
	}
	return maxSum
}
