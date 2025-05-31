package medium

// Find all the possible sums that u can make from the array using a SET
// Then check if it is possible to make sum/2
func canPartition(nums []int) bool {

	sm := 0
	for _, n := range nums {
		sm += n
	}
	if sm%2 != 0 {
		return false
	}

	set := map[int]int{}
	set[0] = 1

	for _, n := range nums {
		newSet := map[int]int{}
		for k := range set {
			newSet[k+n] = 1
			newSet[k] = 1
		}
		set = newSet
	}

	return set[sm/2] == 1

}
