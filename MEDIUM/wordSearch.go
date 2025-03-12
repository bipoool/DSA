package medium

func exist(board [][]byte, word string) bool {
	isVisited := make([][]int, len(board))
	for i := range len(isVisited) {
		isVisited[i] = make([]int, len(board[0]))
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				res := helperExist(board, word, isVisited, 0, i, j)
				if res {
					return true
				}
			}
		}
	}
	return false
}

func helperExist(board [][]byte, target string, isVisited [][]int, currLen int, x int, y int) bool {

	if x >= len(board) || y >= len(board[0]) || x < 0 || y < 0 {
		return false
	}

	if isVisited[x][y] == 1 {
		return false
	}

	if board[x][y] != target[currLen] {
		return false
	}
	if currLen == len(target)-1 {
		return true
	}

	isVisited[x][y] = 1
	v1 := helperExist(board, target, isVisited, currLen+1, x-1, y)
	v2 := helperExist(board, target, isVisited, currLen+1, x+1, y)
	v3 := helperExist(board, target, isVisited, currLen+1, x, y-1)
	v4 := helperExist(board, target, isVisited, currLen+1, x, y+1)
	if v1 || v2 || v3 || v4 {
		return true
	}

	isVisited[x][y] = 0
	return false
}

// func main() {

// 	board := [][]byte{
// 		[]byte{'A', 'B', 'C', 'E'},
// 		[]byte{'S', 'F', 'C', 'S'},
// 		[]byte{'A', 'D', 'E', 'E'},
// 	}

// 	as := exist(board, "ABCB")
// 	println(as)
// }
