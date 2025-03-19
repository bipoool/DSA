package medium

import (
	"sort"
)

func merge(intervals [][]int) [][]int {

	res := [][]int{}
	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	for i := range len(intervals) - 1 {
		if intervals[i][1] >= intervals[i+1][0] {
			intervals[i+1] = []int{min(intervals[i][0], intervals[i+1][0]), max(intervals[i][1], intervals[i+1][1])}
		} else {
			res = append(res, intervals[i])
		}
	}

	res = append(res, intervals[len(intervals)-1])

	return res
}
