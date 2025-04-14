package medium

func topKFrequent(nums []int, k int) []int {

	hmap := map[int]int{}
	bucket := make([][]int, len(nums)+1)

	for i := range nums {
		hmap[nums[i]]++
	}

	for k, v := range hmap {
		if val := bucket[v]; val == nil {
			bucket[v] = []int{}
		}
		bucket[v] = append(bucket[v], k)
	}

	res := []int{}

	for i := len(bucket) - 1; i >= 0; i++ {

		if bucket[i] == nil {
			continue
		}

		if len(bucket[i]) > k {
			res = append(res, bucket[i][:k]...)
		} else {
			res = append(res, bucket[i]...)
		}

	}
	return res

}
