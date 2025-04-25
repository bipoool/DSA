package medium

import (
	"sort"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func getOrder(tasks [][]int) []int {

	for i := range tasks {
		tasks[i] = append(tasks[i], i)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][0] < tasks[j][0]
	})

	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		if a.([]int)[1] == b.([]int)[1] {
			return a.([]int)[2] - b.([]int)[2]
		}
		return a.([]int)[1] - b.([]int)[1]
	})

	res := []int{}
	time := tasks[0][0]
	i := 0

	for {
		if i < len(tasks) && pq.Size() == 0 && tasks[i][0] > time {
			time = tasks[i][0]
		}
		for i < len(tasks) && tasks[i][0] <= time {
			pq.Enqueue(tasks[i])
			i++
		}

		if pq.Size() == 0 && i >= len(tasks) {
			break
		}
		if pq.Size() != 0 {
			task, _ := pq.Dequeue()
			res = append(res, task.([]int)[2])
			time += task.([]int)[1]
		}
	}
	return res
}
