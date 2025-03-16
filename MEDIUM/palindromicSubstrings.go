package medium

func countSubstrings(s string) int {

	result := 0
	for i := range s {
		res1 := expandPal(s, i, i)
		res2 := expandPal(s, i, i+1)
		result += res1 + res2
	}
	return result
}

func expandPal(s string, x int, y int) int {

	res := 0
	for x >= 0 && y < len(s) && s[x] == s[y] {
		res++
		x--
		y++
	}

	return res

}
