package medium

import "container/heap"

type PQ [][]int

func (pq PQ) Len() int {
	return len(pq)
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq PQ) Less(i, j int) bool {
	return pq[i][0]*pq[i][0]+pq[i][1]*pq[i][1] < pq[j][0]*pq[j][0]+pq[j][1]*pq[j][1]
}

func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.([]int))
}

func (pq *PQ) Pop() any {
	val := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return val
}

func kClosest(points [][]int, k int) [][]int {

	prq := PQ{}
	heap.Init(&prq)

	for i := range points {
		point := points[i]
		heap.Push(&prq, point)
		if prq.Len() > k {
			prq.Pop()
		}
	}
	return prq

}
