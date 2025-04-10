package medium

func partitionLabels(s string) []int {

	lastMap := map[byte]int{}
	for i := range s {
		lastMap[s[i]] = i
	}

	res := []int{}
	l := 0
	r := lastMap[s[l]]

	count := 0

	for r <= len(s) {
		for l != r {
			r = max(r, lastMap[s[l]])
			l++
			count++
		}
		res = append(res, count)
		r++
		count = 0
	}
	return res

}
