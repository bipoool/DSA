package medium

var dp [][]int

func findTargetSumWays(nums []int, target int) int {

	dp = make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2)
		dp[i][0] = -1
		dp[i][1] = -1
	}
	a := dfsFindTargetSumWays(nums, target, 0, 0, 0) + dfsFindTargetSumWays(nums, target, 0, 0, 1)
	return a
}

func dfsFindTargetSumWays(nums []int, target int, currSum int, currI int, op int) int {

	if target == currSum && currI == len(nums) {
		return 1
	}

	if currI >= len(nums) {
		return 0
	}

	if dp[currI][op] != -1 {
		return dp[currI][op]
	}

	if op == 0 {
		currSum = currSum + nums[currI]
	} else {
		currSum = currSum - nums[currI]
	}

	dp[currI][op] = dfsFindTargetSumWays(nums, target, currSum, currI+1, 0) + dfsFindTargetSumWays(nums, target, currSum, currI+1, 1)

	return dp[currI][op]

}
