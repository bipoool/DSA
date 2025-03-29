package easy

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
	val := x.(int)
	*pq = append(*pq, val)
}
func (pq *PQ) Pop() any {
	val := (*pq)[len(*pq)-1]
	(*pq) = (*pq)[:len(*pq)-1]
	return val
}

type KthLargest struct {
	prq PQ
	k   int
}

func Constructor(k int, nums []int) KthLargest {

	kth := KthLargest{
		prq: PQ{},
		k:   k,
	}
	heap.Init(&kth.prq)
	for i := 0; i < len(nums); i++ {
		heap.Push(&kth.prq, nums[i])
		if kth.prq.Len() > k {
			heap.Pop(&kth.prq)
		}
	}
	return kth

}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.prq, val)
	if this.prq.Len() > this.k {
		heap.Pop(&this.prq)
	}
	return this.prq[0]
}
