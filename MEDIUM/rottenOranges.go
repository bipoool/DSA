package medium

import "math"

var maxT int
var memO [][]int

func orangesRotting(grid [][]int) int {

	memO = make([][]int, len(grid))
	for i := range memO {
		memO[i] = make([]int, len(grid[0]))
		for j := range memO[i] {
			memO[i][j] = math.MaxInt
		}
	}

	for i := range grid {
		for j := range grid[i] {
			dfso(grid, i, j, 0)
		}
	}
	return maxT
}

func dfso(grid [][]int, x int, y int, ot int) {

	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 {
		return
	}

	if ot+1 >= memO[x][y] {
		return
	}

	memO[x][y] = ot + 1
	maxT = max(maxT, memO[x][y])

	dfso(grid, x+1, y, memO[x][y])
	dfso(grid, x-1, y, memO[x][y])
	dfso(grid, x, y+1, memO[x][y])
	dfso(grid, x, y-1, memO[x][y])

}
