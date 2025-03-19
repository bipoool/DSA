package medium

func insert(intervals [][]int, newInterval []int) [][]int {

	res := [][]int{}

	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > newInterval[1] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			return res
		} else if intervals[i][1] < newInterval[0] {
			res = append(res, intervals[i])
		} else {
			newInterval = []int{min(intervals[i][0], newInterval[0]), max(intervals[i][1], newInterval[1])}
		}

	}
	res = append(res, newInterval)
	return res
}
