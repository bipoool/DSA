package medium

func uniquePaths(m int, n int) int {

	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
	}
	mem[0][0] = 1
	for i := range m {
		for j := range n {
			if i == 0 && j == 0 {
				continue
			}
			up := 0
			down := 0
			if i > 0 {
				up = mem[i-1][j]
			}
			if j > 0 {
				down = mem[i][j-1]
			}
			mem[i][j] = up + down
		}
	}

	return mem[m-1][n-1]

}
