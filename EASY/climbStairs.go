package easy

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	uni := make([]int, n)
	uni[0] = 1
	uni[1] = 2
	for i := 2; i < n; i++ {
		uni[i] = uni[i-1] + uni[i-2]
	}
	return uni[n-1]
}
