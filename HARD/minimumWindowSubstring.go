package hard

import "math"

func minWindow(s string, t string) string {

	if len(s) < len(t) {
		return ""
	}

	tMap := map[byte]int{}
	l := 0
	r := 0
	minLen := math.MaxInt
	count := 0
	startIndex := 0

	for i := range t {
		tMap[t[i]]++
	}

	for r < len(s) {

		if v, ok := tMap[s[r]]; ok {
			tMap[s[r]]--
			if v > 0 {
				count++
			}

		}

		for count == len(t) {
			if r-l < minLen {
				startIndex = l
				minLen = r - l
			}
			if v, ok := tMap[s[l]]; ok {
				tMap[s[l]]++
				if v == 0 {
					count--
				}
			}
			l++
		}
		r++
	}
	return s[startIndex : startIndex+minLen+1]
}
func main() {
	println(minWindow("ADOBECODEBANC", "ABC"))
}
