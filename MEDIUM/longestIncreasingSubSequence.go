package medium

import "slices"

func lengthOfLISDp(nums []int) int {

	mem := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {

		for j := i + 1; j < len(nums); j++ {

			if nums[i] < nums[j] {
				mem[i] = max(mem[j]+1, mem[i])
			}

		}

	}

	return slices.Max(mem)

}

// Creat an empty array
// Loop into nums and try to find out the position at which you should insert the current element
// Find the nearest smallest element to the current element
// And Insert/update it to the next index
func lengthOfLIS(nums []int) int {

	mem := make([]int, 0)

	for i := range len(nums) {

		l := 0
		r := len(mem) - 1
		target := nums[i]
		for l < r {
			mid := (l + r) / 2
			if mem[mid] == target {
				l = mid
				break
			}
			if mem[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		if l < len(mem) && mem[l] >= target {
			l = l - 1
		}

		if l+1 < len(mem) {
			mem[l+1] = target
		} else {
			mem = append(mem, target)
		}
	}

	return len(mem)

}
