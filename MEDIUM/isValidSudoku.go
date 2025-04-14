package medium

func isValidSudoku(board [][]byte) bool {

	rowMap := [9]map[byte]int{}
	colMap := [9]map[byte]int{}

	gridMap := map[string]map[byte]int{}

	for i := range board {
		for j := range board[i] {
			if val := rowMap[i]; val == nil {
				rowMap[i] = map[byte]int{}
			}
			if val := colMap[j]; val == nil {
				colMap[j] = map[byte]int{}
			}
			rowMap[i][board[i][j]]++
			colMap[j][board[i][j]]++
			gridKey := string(i%3) + ":" + string(j%3)
			if _, ok := gridMap[gridKey]; !ok {
				gridMap[gridKey] = map[byte]int{}
			}
			gridMap[gridKey][board[i][j]]++
		}
	}

	for i := range board {
		for j := range board[i] {
			if rowMap[i][board[i][j]] > 1 {
				return false
			}
			if colMap[j][board[i][j]] > 1 {
				return false
			}

			gridKey := string(i%3) + ":" + string(j%3)
			if gridMap[gridKey][board[i][j]] > 1 {
				return false
			}
		}
	}
	return true
}
