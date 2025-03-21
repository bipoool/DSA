package medium

func longestConsecutive(nums []int) int {

	hmap := map[int]int{}

	for i := range nums {
		hmap[nums[i]] = 1
	}

	count := 0
	maxCount := 0

	for i := range nums {
		ele := nums[i]
		if hmap[ele-1] != 0 {
			continue
		}
		for hmap[ele] != 0 {
			count++
			ele++
		}

		maxCount = max(count, maxCount)
		count = 0

	}
	return maxCount
}
