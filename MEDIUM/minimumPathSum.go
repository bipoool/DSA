package medium

// Loop from top to bottom
// Check for each index = calculate the sum possible from top and bottom
// Only 2 ways possible -> from top and from right -> take minimum
func minPathSum(grid [][]int) int {
	for i := range grid {
		for j := range grid[0] {
			if i > 0 && j > 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			} else if i > 0 {
				grid[i][j] += grid[i-1][j]
			} else if j > 0 {
				grid[i][j] += grid[i][j-1]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
