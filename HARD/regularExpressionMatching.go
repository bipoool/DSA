package hard

var dpMatch [][]int

func isMatch(s string, p string) bool {

	dpMatch = make([][]int, len(s)+1)
	for i := range dpMatch {
		dpMatch[i] = make([]int, len(p)+1)
	}
	dpMatch[len(s)][len(p)] = 1

	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			match := i < len(s) && (s[i] == p[j] || p[j] == '.')

			if j+1 < len(p) && p[j+1] == '*' {
				dpMatch[i][j] = dpMatch[i][j+2]
				if match {
					if dpMatch[i][j] == 1 || dpMatch[i+1][j] == 1 {
						dpMatch[i][j] = 1
					}
				}
			} else if match {
				dpMatch[i][j] = dpMatch[i+1][j+1]
			}
		}
	}
	return dpMatch[0][0] == 1

}
