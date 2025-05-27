package medium

import "math"

// Expand accross all the centers -> i & i+1
// Find the length of the palindrom at each center
func longestPalindrome(s string) string {

	maxLen := math.MinInt
	st := 0

	for i := 0; i < len(s); i++ {

		l1, r1 := expand(s, i, i)
		l2, r2 := expand(s, i, i+1)
		if r1-l1+1 > maxLen {
			maxLen = r1 - l1 + 1
			st = l1
		}
		if r2-l2+1 > maxLen {
			maxLen = r2 - l2 + 1
			st = l2
		}

	}

	return s[st : st+maxLen]

}

func expand(s string, l int, r int) (int, int) {
	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			break
		}
		l--
		r++
	}
	l++
	r--
	return l, r
}
