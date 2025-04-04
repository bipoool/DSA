package medium

var dpint [][]int

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dpint = make([][]int, len(s1))
	for i := range dpint {
		dpint[i] = make([]int, len(s2))
	}
	return dfsIsInterleave(s3, s1, s2, "", 0, 0)
}

func dfsIsInterleave(target string, s1 string, s2 string, currStr string, s1i int, s2i int) bool {

	if target == currStr && len(s1+s2) == 0 {
		return true
	}

	if len(s1) == 0 || len(s2) == 0 {
		return target == currStr+s1+s2
	}

	if dpint[s1i][s2i] != 0 {
		return dpint[s1i][s2i] == 1
	}

	if s1[0] == target[s1i+s2i] && dfsIsInterleave(target, s1[1:], s2, currStr+string(s1[0]), s1i+1, s2i) ||
		s2[0] == target[s1i+s2i] && dfsIsInterleave(target, s1, s2[1:], currStr+string(s2[0]), s1i, s2i+1) {
		dpint[s1i][s2i] = 1
	} else {
		dpint[s1i][s2i] = -1
	}
	return dpint[s1i][s2i] == 1
}
