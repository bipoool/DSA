package medium

func dailyTemperatures(temperatures []int) []int {

	stack := make([][]int, len(temperatures))
	start := -1
	stack[start+1] = []int{temperatures[len(temperatures)-1], len(temperatures) - 1}
	start++

	res := make([]int, len(temperatures))

	for i := len(temperatures) - 2; i >= 0; i-- {

		for start >= 0 && stack[start][0] <= temperatures[i] {
			start--
		}
		if start == -1 {
			res[i] = 0
		} else {
			res[i] = stack[start][1] - i
		}
		stack[start+1] = []int{temperatures[i], i}
		start++
	}
	return res
}
