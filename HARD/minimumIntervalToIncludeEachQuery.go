package hard

import (
	"sort"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

func minInterval(intervals [][]int, queries []int) []int {

	cmp := func(a, b interface{}) int {
		return a.([2]int)[0] - b.([2]int)[0]
	}

	pq := priorityqueue.NewWith(cmp)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	queryNew := make([][]int, 0)

	for i, v := range queries {
		queryNew = append(queryNew, []int{i, v})
	}

	sort.Slice(queryNew, func(i, j int) bool {
		return queryNew[i][1] < queryNew[j][1]
	})

	res := make([]int, len(queries))
	inti := 0
	for i := range queryNew {

		for inti < len(intervals) {
			if intervals[inti][0] <= queryNew[i][1] {
				l := intervals[inti][0]
				r := intervals[inti][1]
				pq.Enqueue([2]int{r - l + 1, r})
				inti++
			} else {
				break
			}
		}

		for !pq.Empty() {
			crrEl, _ := pq.Peek()
			if queryNew[i][1] > crrEl.([2]int)[1] {
				pq.Dequeue()
			} else {
				break
			}
		}

		len := -1
		if !pq.Empty() {
			crrEl, _ := pq.Peek()
			len = crrEl.([2]int)[0]
		}
		res[queryNew[i][0]] = len

	}

	return res
}
