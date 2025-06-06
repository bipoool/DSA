package medium

// Question - You are given a 0-indexed integer array stones sorted in strictly increasing order representing the positions of stones in a river.
// A frog, initially on the first stone, wants to travel to the last stone and then return to the first stone. However, it can jump to any stone at most once.
// The length of a jump is the absolute difference between the position of the stone the frog is currently on and the position of the stone to which the frog jumps.
// More formally, if the frog is at stones[i] and is jumping to stones[j], the length of the jump is |stones[i] - stones[j]|.
// The cost of a path is the maximum length of a jump among all jumps in the path.

// Just calculate Odd index paths and even index paths
// Calculate the maximum difference you encounter
// Why even and odd, if you take 2 steps the difference would be always higher because the array is sorted
func maxJump(stones []int) int {

	if len(stones) == 2 {
		return stones[1] - stones[0]
	}

	res := stones[0]
	for i := 2; i < len(stones); i += 2 {
		res = max(res, stones[i]-stones[i-2])
	}
	for i := 3; i < len(stones); i += 2 {
		res = max(res, stones[i]-stones[i-2])
	}
	return res

}
