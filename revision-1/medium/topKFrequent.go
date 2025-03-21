package medium

func topKFrequent(nums []int, k int) []int {
	hmap := map[int]int{}

	for i := range nums {
		hmap[nums[i]] += 1
	}
	revMap := map[int][]int{}

	for i := range hmap {
		if _, ok := revMap[hmap[i]]; !ok {
			revMap[hmap[i]] = make([]int, 0)
		}
		revMap[hmap[i]] = append(revMap[hmap[i]], i)
	}
	arr := make([][]int, len(nums)+1)

	for i := range revMap {
		arr[i] = revMap[i]
	}
	res := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != nil {
			elemetnts := arr[i]
			lenArr := len(arr[i])

			if lenArr > k {
				res = append(res, elemetnts[:k]...)
			} else {
				res = append(res, arr[i]...)
			}

			k -= lenArr
		}
		if k <= 0 {
			break
		}
	}
	return res
}
