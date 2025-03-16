package medium

import "math"

func longestPalindrome(s string) string {

	maxLen := math.MinInt
	startId := 0
	for i := range s {
		odd, startOdd := expand(s, i, i)
		even, startEven := expand(s, i, i+1)

		if odd > even && maxLen < odd {
			maxLen = odd
			startId = startOdd
		} else if maxLen < even {
			maxLen = even
			startId = startEven
		}
	}
	return s[startId : startId+maxLen]
}

func expand(s string, x int, y int) (int, int) {

	for x >= 0 && y < len(s)-1 && s[x] == s[y] {
		x--
		y++
	}

	if x >= 0 && y < len(s)-1 && s[x] == s[y] {

	}
	return y - x - 1, x + 1

}

// func main() {
// 	println(longestPalindrome("bb"))
// }
