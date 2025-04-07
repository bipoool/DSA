package medium

import "math"

func jump(nums []int) int {

	l, r := 0, 0
	far := math.MinInt
	res := 0
	for r < len(nums) {

		for i := l; i <= r; i++ {
			far = max(far, nums[i]+i)
		}
		l = r + 1
		r = far
		res++
	}
	return res

}
