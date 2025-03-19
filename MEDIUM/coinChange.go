package medium

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 1; i < amount; i++ {
		for j := range coins {
			if i-coins[j] >= 0 {
				dp[i] = min(dp[i], 1+dp[i-coins[j]])
			}
		}
	}
	return dp[amount]
}
