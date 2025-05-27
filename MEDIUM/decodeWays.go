package medium

// DP questions
// Create and array with len(s) + 1 elements
// if cur > 0 then mem[i+1] = mem[i] because it was supposed to be considered as a separate ele
// If combined <= 26 || >= 10 then add mem[i-1] because then you are invalidating previos case of
// mem[i] & adding the case before it

func numDecodings(s string) int {

	if s[0] == '0' {
		return 0
	}
	mem := make([]int, len(s)+1)
	mem[0] = 1
	mem[1] = 1

	for i := 1; i < len(s); i++ {

		cur := int(s[i] - '0')
		prev := int(s[i-1] - '0')

		comb := (prev * 10) + cur

		if cur > 0 {
			mem[i+1] = mem[i]
		}
		if comb < 27 && comb >= 10 {
			mem[i+1] += mem[i-1]
		}

	}
	return mem[len(s)]

}
