package medium

func pacificAtlantic(heights [][]int) [][]int {

	visited := make([][]int, len(heights))
	for i := range visited {
		visited[i] = make([]int, len(heights[i]))
	}

	res := make([][]int, 0)

	for i := range heights {
		for j := range heights[i] {
			if visited[i][j] == 2 {
				continue
			}
			a1, a2 := helper3(heights, i, j, heights[i][j], &visited, &res)
			if a1 && a2 {
				res = append(res, []int{i, j})
				visited[i][j] = 2
			}
		}
	}
	return res

}

func helper3(heights [][]int, x int, y int, parent int, visited *[][]int, res *[][]int) (bool, bool) {

	if x < 0 || y < 0 || x >= len(heights) || y >= len(heights[0]) {
		return false, false
	}

	if (*visited)[x][y] == 1 {
		return false, false
	}

	if (*visited)[x][y] == 2 && heights[x][y] <= parent {
		return true, true
	}

	if heights[x][y] > parent {
		return false, false
	}

	pf := false
	af := false
	if x == 0 || y == 0 {
		pf = true
	}
	if x == len(heights)-1 || y == len(heights[0])-1 {
		af = true
	}

	if pf && af {
		(*visited)[x][y] = 2
		*res = append(*res, []int{x, y})
		return pf, af
	}

	(*visited)[x][y] = 1
	pf1, af1 := helper3(heights, x-1, y, heights[x][y], visited, res)
	pf2, af2 := helper3(heights, x+1, y, heights[x][y], visited, res)
	pf3, af3 := helper3(heights, x, y-1, heights[x][y], visited, res)
	pf4, af4 := helper3(heights, x, y+1, heights[x][y], visited, res)

	if (*visited)[x][y] != 2 && ((pf || pf1 || pf2 || pf3 || pf4) && (af || af1 || af2 || af3 || af4)) {
		(*visited)[x][y] = 2
		*res = append(*res, []int{x, y})
	} else {
		(*visited)[x][y] = 0
	}

	return pf || pf1 || pf2 || pf3 || pf4, af || af1 || af2 || af3 || af4

}
