package medium

import "container/heap"

type PQ []int

func (pq PQ) Len() int {
	return len(pq)
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq PQ) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.(int))
}
func (pq *PQ) Pop() any {
	val := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return val
}

func findKthLargest(nums []int, k int) int {

	pq := PQ{}
	for i := range nums {
		heap.Push(&pq, nums[i])
		if pq.Len() > k {
			heap.Pop(&pq)
		}
	}
	heap.Init(&pq)
	return heap.Pop(&pq).(int)

}
