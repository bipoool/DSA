package medium

import "math"

func characterReplacement(s string, k int) int {

	l := 0
	r := 0

	hMap := map[byte]int{}
	maxLen := math.MinInt

	for r < len(s) {

		hMap[s[r]]++
		currLen := r - l + 1

		if currLen-hMap[s[r]] > k {
			l++
		}

		if maxLen < r-l+1 {
			maxLen = r - l + 1
		}

		r++
	}

	return maxLen

}
