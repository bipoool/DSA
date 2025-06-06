package medium

// Loop from top to bottom
// Check for each index = number of ways to land on that point
// Only 2 ways possible -> from top and from right -> add those ways
func uniquePathsWithObstacles(o [][]int) int {
	if o[len(o)-1][len(o[0])-1] == 1 || o[0][0] == 1 {
		return 0
	}
	o[0][0] = 1
	for i := 0; i < len(o); i++ {
		for j := 0; j < len(o[0]); j++ {
			if o[i][j] == 1 && !(i == 0 && j == 0) {
				o[i][j] = -1
				continue
			}
			if i > 0 && o[i-1][j] != -1 {
				o[i][j] += o[i-1][j]
			}
			if j > 0 && o[i][j-1] != -1 {
				o[i][j] += o[i][j-1]
			}
		}
	}

	return o[len(o)-1][len(o[0])-1]
}
