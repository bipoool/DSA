package easy

import "container/heap"

func lastStoneWeight(stones []int) int {
	prq := PQ{}
	heap.Init(&prq)
	for i := range stones {
		heap.Push(&prq, stones[i])
	}
	for prq.Len() > 1 {
		el1 := heap.Pop(&prq)
		el2 := heap.Pop(&prq)
		heap.Push(&prq, el1.(int)-el2.(int))
	}
	return heap.Pop(&prq).(int)
}
