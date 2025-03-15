package medium

type trie struct {
	trieMap map[byte]*trie
	isWord  bool
}

func findWords(board [][]byte, words []string) []string {
	rootTrie := convertToTrie(words)
	visited := make([][]int, len(board))

	for i := range board {
		visited[i] = make([]int, len(board[0]))
	}

	res := []string{}

	for i := range board {
		for j := range board[i] {
			helper1(board, &res, visited, "", rootTrie, i, j)
		}
	}

	return res

}

func helper1(board [][]byte, res *[]string, visited [][]int, currstr string, currTrie *trie, x int, y int) {

	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) {
		return
	}

	if visited[x][y] == 1 {
		return
	}

	v, ok := currTrie.trieMap[board[x][y]]
	if ok && v.isWord {
		*res = append(*res, currstr+string(board[x][y]))
		v.isWord = false
	}

	if !ok {
		return
	}

	currTrie = currTrie.trieMap[board[x][y]]
	visited[x][y] = 1
	helper1(board, res, visited, currstr+string(board[x][y]), currTrie, x+1, y)
	helper1(board, res, visited, currstr+string(board[x][y]), currTrie, x-1, y)
	helper1(board, res, visited, currstr+string(board[x][y]), currTrie, x, y+1)
	helper1(board, res, visited, currstr+string(board[x][y]), currTrie, x, y-1)
	visited[x][y] = 0
}

func convertToTrie(words []string) *trie {

	rootTrie := &trie{
		trieMap: map[byte]*trie{},
	}

	for _, word := range words {
		currTrie := rootTrie
		for al := range word {
			if tr, ok := currTrie.trieMap[word[al]]; ok {
				currTrie = tr
			} else {
				currTrie.trieMap[word[al]] = &trie{
					trieMap: map[byte]*trie{},
				}
				currTrie = currTrie.trieMap[word[al]]
			}
		}
		currTrie.isWord = true
	}

	return rootTrie
}

// func main() {
// 	board := [][]byte{
// 		[]byte{'a', 'b', 'c'},
// 		[]byte{'a', 'e', 'd'},
// 		[]byte{'a', 'f', 'g'},
// 	}

// 	tri := findWords(board, []string{"abcdefg"})

// 	println(tri)
// 	// for i := range tri {
// 	// 	println(tri[i])
// 	// }
// }
