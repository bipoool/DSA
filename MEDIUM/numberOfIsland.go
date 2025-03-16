package medium

func numIslands(grid [][]byte) int {

	visited := make([][]int, len(grid))
	for i := range visited {
		visited[i] = make([]int, len(grid[i]))
	}

	res := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && visited[i][j] != 1 {
				helper4(grid, &visited, i, j)
				res++
			}
		}
	}
	return res

}

func helper4(grid [][]byte, visited *[][]int, x int, y int) bool {

	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return false
	}

	if (*visited)[x][y] == 1 {
		return false
	}

	if grid[x][y] == '0' {
		return false
	}
	(*visited)[x][y] = 1
	helper4(grid, visited, x-1, y)
	helper4(grid, visited, x+1, y)
	helper4(grid, visited, x, y-1)
	helper4(grid, visited, x, y+1)

	return true

}
