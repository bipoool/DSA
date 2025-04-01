package hard

var resq [][]string

func solveNQueens(n int) [][]string {

	resq = make([][]string, 0)
	helperQ([]string{}, getNewRow(n), 0)
	return resq

}

func helperQ(b []string, currRow string, x int) {

	if x == len(currRow) {
		temp := make([]string, len(b))
		copy(temp, b)
		resq = append(resq, temp)
		return
	}

	for j := 0; j < len(currRow); j++ {

		if isSafe(b, x, j) {
			newRow := currRow[0:j] + "Q"
			if j+1 < len(currRow) {
				newRow = newRow + currRow[j+1:]
			}
			b = append(b, newRow)
			helperQ(b, getNewRow(len(currRow)), x+1)
			newRow = currRow[0:j] + "."
			if j+1 < len(currRow) {
				newRow = newRow + currRow[j+1:]
			}
			b = b[:len(b)-1]
		}

	}

}

func isSafe(b []string, x int, y int) bool {

	// Check Vertically
	for i := 0; i < len(b) && y < len(b[i]); i++ {
		if b[i][y] == 'Q' {
			return false
		}
	}

	// Check left-up diagonal
	i := x
	j := y

	for i-1 >= 0 && j-1 >= 0 {
		if b[i-1][j-1] == 'Q' {
			return false
		}
		i--
		j--
	}

	// Check right-up diagonal
	i = x
	j = y

	for i-1 >= 0 && j+1 < len(b[i-1]) {
		if b[i-1][j+1] == 'Q' {
			return false
		}
		i--
		j++
	}

	// Check left-down diagonal
	i = x
	j = y

	for i+1 < len(b) && j-1 >= 0 {
		if b[i+1][j-1] == 'Q' {
			return false
		}
		i++
		j--
	}

	// Check right-down diagonal
	i = x
	j = y

	for i+1 < len(b) && j+1 < len(b[i+1]) {
		if b[i+1][j+1] == 'Q' {
			return false
		}
		i++
		j++
	}
	return true
}

func getNewRow(n int) string {
	var newRow2 string
	for i := 0; i < n; i++ {
		newRow2 = newRow2 + "."
	}
	return newRow2
}
