package hard

var dp [][]int

func longestIncreasingPath(matrix [][]int) int {

	dp = make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			dfs(matrix, i, j, -1, 0)
		}
	}

	for i := range dp {
		for j := range dp[i] {
			dp[0][0] = max(dp[0][0], dp[i][j])
		}
	}
	return dp[0][0]

}

func dfs(m [][]int, x int, y int, pval int, currLen int) int {

	if x < 0 || y < 0 || x >= len(m) || y >= len(m[0]) {
		return currLen
	}

	if m[x][y] <= pval {
		return currLen
	}

	if dp[x][y] != -1 {
		return dp[x][y]
	}

	dp[x][y] = max(dfs(m, x+1, y, m[x][y], currLen)+1, dp[x][y])
	dp[x][y] = max(dfs(m, x-1, y, m[x][y], currLen)+1, dp[x][y])
	dp[x][y] = max(dfs(m, x, y+1, m[x][y], currLen)+1, dp[x][y])
	dp[x][y] = max(dfs(m, x, y-1, m[x][y], currLen)+1, dp[x][y])

	return dp[x][y]
}
