package medium

import "strconv"

func numDecodings(s string) int {

	mem := make([]int, len(s)+1)
	mem[0] = 1
	mem[1] = 1

	for i := 1; i < len(s)-1; i++ {
		currDig, _ := strconv.ParseInt(string(s[i]), 10, 0)
		nextDig, _ := strconv.ParseInt(string(s[i+1]), 10, 0)

		combined := currDig*10 + nextDig

		if combined <= 26 {
			mem[i] = mem[i-1] * 2
		} else {
			mem[i] = mem[i-1]
		}
	}

	return mem[len(s)-1]
}
