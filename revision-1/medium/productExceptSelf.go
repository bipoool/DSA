package medium

func productExceptSelf(nums []int) []int {

	prod := 1
	zeroCount := 0

	for i := range nums {
		if nums[i] == 0 {
			zeroCount++
			continue
		}
		prod *= nums[i]
	}

	for i := range nums {
		if zeroCount > 1 {
			nums[i] = 0
			continue
		}
		if nums[i] != 0 && zeroCount > 0 {
			nums[i] = 0
			continue
		}
		if zeroCount > 0 {
			nums[i] = prod
		} else {
			nums[i] = prod / nums[i]
		}
	}
	return nums
}
