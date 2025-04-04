package medium

var dpmem [][]int

func solve(board [][]byte) {
	dpmem = make([][]int, len(board))
	for i := range dpmem {
		dpmem[i] = make([]int, len(board[i]))
	}
	for i := range board {
		for j := range board[i] {

			if board[i][j] == 'O' && dpmem[i][j] == 0 {
				helperSur(board, i, j)
			}

		}
	}
}

func helperSur(board [][]byte, x int, y int) bool {

	if x < 0 || x >= len(board) || y < 0 || y >= len(board) {
		return false
	}

	if board[x][y] == 'X' {
		return true
	}

	if dpmem[x][y] == 2 {
		return false
	}
	if dpmem[x][y] == 1 {
		return true
	}

	board[x][y] = 'X'

	a1 := helperSur(board, x-1, y)
	a2 := helperSur(board, x+1, y)
	a3 := helperSur(board, x, y-1)
	a4 := helperSur(board, x, y+1)

	if a1 && a2 && a3 && a4 {
		dpmem[x][y] = 1
	} else {
		dpmem[x][y] = 2
		board[x][y] = 'O'
	}

	return a1 && a2 && a3 && a4

}

// func main() {

// 	board := [][]byte{
// 		[]byte{'X', 'O', 'X', 'O', 'X', 'O'},
// 		[]byte{'O', 'X', 'O', 'X', 'O', 'X'},
// 		[]byte{'X', 'O', 'X', 'O', 'X', 'O'},
// 		[]byte{'O', 'X', 'O', 'X', 'O', 'X'},
// 	}
// 	solve(board)

// }
