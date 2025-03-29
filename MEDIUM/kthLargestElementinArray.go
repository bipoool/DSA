package medium

import "container/heap"

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
