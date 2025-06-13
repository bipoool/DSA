package medium

// Each dp1[i][j] represents the if j can be made either using i or not using i
var dp1 [][]bool

func canPartition(nums []int) bool {
	sm := 0
	dp1 = make([][]bool, len(nums))
	for i := range nums {
		sm += nums[i]
	}
	if sm%2 != 0 {
		return false
	}
	sm /= 2

	for i := range dp {
		dp1[i] = make([]bool, sm+1)
		dp1[i][0] = true
	}

	for id := len(nums) - 2; id >= 0; id-- {
		for t := 1; t <= sm; t++ {
			notTake := dp1[id+1][t]
			take := false
			if nums[id] <= t {
				take = dp1[id+1][t-nums[id]]
			}
			dp1[id][t] = take || notTake
		}
	}

	return dp1[0][sm]
}
