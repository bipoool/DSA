package easy

// At each index
// Minimum cost is min(nums[i-1], nums[i-2]) + nums[i]
// Return min(cost[len(cost)-1], cost[len(cost)-2])
func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}
	return min(cost[len(cost)-1], cost[len(cost)-2])
}
