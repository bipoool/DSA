package medium

var memdis [][]int

func minDistance(word1 string, word2 string) int {

	memdis = make([][]int, len(word1)+1)
	for i := range memdis {
		memdis[i] = make([]int, len(word2)+1)
		memdis[i][len(word2)] = len(word1) - i
	}
	for i := range memdis[0] {
		memdis[len(word1)][i] = len(word2) - i
	}
	memdis[len(word1)][len(word2)] = 0
	for i := len(word1) - 1; i >= 0; i-- {
		for j := len(word2) - 1; j >= 0; j-- {

			if word1[i] == word2[j] {
				memdis[i][j] = memdis[i+1][j+1]
			} else {
				memdis[i][j] = min(memdis[i+1][j], memdis[i][j+1], memdis[i+1][j+1]) + 1
			}

		}
	}
	return memdis[0][0]

}
