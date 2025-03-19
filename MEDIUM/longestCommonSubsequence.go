package medium

func longestCommonSubsequence(text1 string, text2 string) int {

	mem := make([][]int, len(text1)+1)
	for i := range mem {
		mem[i] = make([]int, len(text2)+1)
	}

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				mem[i][j] = mem[i+1][j+1] + 1
			} else {
				mem[i][j] = max(mem[i+1][j], mem[i][j+1])
			}
		}
	}

	return mem[0][0]

}
