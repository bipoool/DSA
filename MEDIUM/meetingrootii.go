package medium

import "sort"

// Check if we can use existing any existing meeting room
// Create two seperate array of start and end
// If start is less than the end => we need new root
// If not => update the end array index

func minMeetingRooms(in [][]int) int {

	st := []int{}
	end := []int{}

	for i := range in {
		st = append(st, in[i][0])
		end = append(end, in[i][1])
	}
	sort.Slice(st, func(i, j int) bool {
		return st[i] < st[j]
	})
	sort.Slice(end, func(i, j int) bool {
		return end[i] < end[j]
	})

	count := 0
	ei := 0
	for i := range st {
		if st[i] < end[ei] {
			count++
		} else {
			ei++
		}
	}
	return count

}
