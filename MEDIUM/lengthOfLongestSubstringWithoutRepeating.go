package medium

import "math"

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}
	l := 0
	r := 0

	hmap := map[byte]int{}
	maxLen := math.MinInt

	for r < len(s) {

		currChar := s[r]
		hmap[s[r]]++

		v, ok := hmap[currChar]
		for ok && v > 1 && l < r {
			hmap[s[l]]--
			l++
			v, ok = hmap[currChar]
		}

		if maxLen < (r - l + 1) {
			maxLen = r - l + 1
		}
		r++

	}

	return maxLen
}
