package medium

// Every time you have 2 option -> Add or subtract
// Create a recusive function to do so
// Then cache the output for every (i, currSum) because these are the only 2 variable
var dp []map[int]int

func findTargetSumWays(nums []int, target int) int {
	dp = make([]map[int]int, len(nums))
	for i := range dp {
		dp[i] = map[int]int{}
	}
	return dfsfindTargetSumWays(nums, 0, 0, target)
}

func dfsfindTargetSumWays(nums []int, id int, currSum int, target int) int {
	if currSum == target && id == len(nums) {
		return 1
	}

	if id == len(nums) {
		return 0
	}

	if dp[id][currSum] != 0 {
		return dp[id][currSum]
	}

	dp[id][currSum] = dfsfindTargetSumWays(nums, id+1, currSum+nums[id], target) + dfsfindTargetSumWays(nums, id+1, currSum-nums[id], target)
	return dp[id][currSum]
}
