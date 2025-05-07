package medium

func canSeePersonsCount(h []int) []int {
	st := []int{}
	ans := make([]int, len(h))
	ans[len(ans)-1] = 0
	for i := len(h) - 1; i >= 0; i-- {
		count := 0
		for len(st) != 0 && st[len(st)-1] < h[i] {
			st = st[:len(st)-1]
			count++
		}
		if len(st) != 0 {
			count++
		}
		ans[i] = count
		st = append(st, h[i])
	}
	return ans
}
