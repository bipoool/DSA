package medium

import "math"

var visited [][]int
var maxAre int

func maxAreOfIsland(grid [][]int) int {

	visited = make([][]int, len(grid))

	maxAre = math.MinInt

	for i := range grid {
		visited[i] = make([]int, len(grid[0]))
	}

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				maxAre = max(helperIsLand(grid, i, j, 0), maxAre)
			}
		}
	}
	return maxAre

}

func helperIsLand(grid [][]int, x int, y int, currArea int) int {

	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 {
		return currArea
	}

	if visited[x][y] == 1 {
		return currArea
	}

	currArea++

	visited[x][y] = 1

	currArea = helperIsLand(grid, x+1, y, currArea)
	currArea = helperIsLand(grid, x-1, y, currArea)
	currArea = helperIsLand(grid, x, y+1, currArea)
	currArea = helperIsLand(grid, x, y-1, currArea)

	return currArea

}
