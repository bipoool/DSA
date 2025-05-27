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

// DFS solution top down
var memcoinChange []int

func coinChangeDfs(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}
	memcoinChange = make([]int, amount+1)

	for i := range mem {
		mem[i] = -1
	}
	memcoinChange[0] = 0
	helpcoinChange(coins, amount)

	if memcoinChange[amount] == math.MaxInt {
		return -1
	}

	return memcoinChange[amount]

}

func helpcoinChange(coins []int, cursum int) int {

	if memcoinChange[cursum] != -1 {
		return memcoinChange[cursum]
	}

	res := math.MaxInt
	for i := range coins {
		if cursum-coins[i] >= 0 {
			r1 := helpcoinChange(coins, cursum-coins[i])
			if r1 == math.MaxInt {
				continue
			}
			res = min(res, 1+r1)
		}
	}

	memcoinChange[cursum] = res
	return res

}
