package medium

import "strconv"

func numDecodings(s string) int {

	if s[0] == '0' {
		return 0
	}
	mem := make([]int, len(s))
	mem[0] = 1
	for i := 1; i < len(s); i++ {
		currDig, _ := strconv.ParseInt(string(s[i]), 10, 0)
		prevDig, _ := strconv.ParseInt(string(s[i-1]), 10, 0)

		combined := prevDig*10 + currDig

		if currDig == 0 {
			if combined <= 26 {
				if i-2 >= 0 {
					mem[i] = mem[i-2]
				} else {
					mem[i] = mem[i-1]
				}
			} else {
				return 0
			}
			continue
		}
		if prevDig == 0 {
			mem[i] = mem[i-1]
			continue
		}

		if combined <= 26 {
			if i-2 < 0 {
				mem[i] = mem[i-1] * 2
			} else {
				mem[i] = mem[i-1] + mem[i-2]
			}
		} else {
			mem[i] = mem[i-1]
		}
	}
	return mem[len(mem)-1]
}
