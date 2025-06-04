package medium

import "math"

// Create 2 arrays
// mx[i] represents the max element till i from left
// mn[i] represents the min element till i from right
// Just compare each element in for loop and add the score
func sumOfBeauties(nums []int) int {
	mx := make([]int, len(nums))
	mn := make([]int, len(nums))

	xe := 0
	ne := math.MaxInt
	for i := 0; i < len(nums); i++ {
		mx[i] = xe
		xe = max(xe, nums[i])
	}

	for i := len(nums) - 1; i >= 0; i-- {
		mn[i] = ne
		ne = min(ne, nums[i])
	}

	sm := 0

	for i := 1; i < len(nums)-1; i++ {
		n := nums[i]
		curmx := mx[i]
		curmn := mn[i]
		prev := nums[i-1]
		next := nums[i+1]

		if n > curmx && n < curmn {
			sm += 2
		} else if n > prev && n < next {
			sm += 1
		}
	}

	return sm
}
