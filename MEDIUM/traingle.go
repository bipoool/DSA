package medium

import "math"

func minimumTotal(t [][]int) int {

	for i := 1; i < len(t); i++ {
		for j := range t[i] {
			if j > 0 && j < len(t[i-1]) {
				t[i][j] += min(t[i-1][j-1], t[i-1][j])
			} else if j > 0 {
				t[i][j] += t[i-1][j-1]
			} else if j < len(t[i-1]) {
				t[i][j] += t[i-1][j]
			}
		}
	}
	res := math.MaxInt

	for i := range t[len(t)-1] {
		res = min(res, t[len(t)-1][i])
	}
	return res

}
