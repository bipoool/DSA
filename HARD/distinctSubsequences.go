package hard

// Create 2D dp array
// Loop in both the strings
// If s[i] == t[j]
func numDistinct(s string, t string) int {

	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
		dp[i][len(t)] = 1
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := len(t) - 1; j >= 0; j-- {
			dp[i][j] = dp[i+1][j]
			if s[i] == t[j] {
				dp[i][j] += dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]

}
