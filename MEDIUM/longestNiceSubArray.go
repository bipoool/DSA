package medium

// Use sliding window
// OR to set the bits
// XOR to unset if curr window is not producing zero
// Why this works?
// If 2 ones overlap, we will have to remove element from the left till they don't overlap
// So when you remove element use XOR, when you find the overlapping element, XOR will produce 1 XOR 1 = 0

func longestNiceSubarray(nums []int) int {

	l := 0
	r := 0
	runAnd := 0
	maxlen := 1

	for r < len(nums) {
		for runAnd&nums[r] != 0 && l < r {
			runAnd ^= nums[l]
			l++
		}
		runAnd = runAnd | nums[r]
		maxlen = max(maxlen, r-l+1)
		r++
	}
	return maxlen

}
