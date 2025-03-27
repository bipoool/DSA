package medium

import "sort"

func carFleet(target int, position []int, speed []int) int {

	sp := make([][]int, len(position))
	for i := range position {
		sp[i] = []int{position[i], speed[i]}
	}

	sort.Slice(sp, func(i, j int) bool {
		return sp[i][0] > sp[j][0]
	})

	stack := make([]float64, 0, len(sp))
	for i := range sp {
		currentTime := float64(target-sp[i][0]) / float64(sp[i][1])
		if len(stack) == 0 || currentTime > stack[len(stack)-1] {
			stack = append(stack, currentTime)
		}
	}
	return len(stack)
}
