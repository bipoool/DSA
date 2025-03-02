package medium

func topKFrequent(nums []int, k int) []int {

	freq := make(map[int]int)
	bucket := make([][]int, len(nums)+1)
	result := make([]int, 0)

	for _, v := range nums {
		freq[v]++
	}

	for k, v := range freq {
		if val := bucket[v]; val == nil {
			bucket[v] = make([]int, 0)
		}
		bucket[v] = append(bucket[v], k)
	}

	for i := len(bucket) - 1; i >= 0 && k > 0; i-- {
		if bucket[i] != nil {
			arr := bucket[i]
			lenArr := len(arr)

			if lenArr <= k {
				result = append(result, bucket[i]...)
			} else {
				result = append(result, bucket[i][0:k]...)
			}
			k -= lenArr
		}
	}
	return result

}
