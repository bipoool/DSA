package medium

import "sort"

func threeSum(nums []int) [][]int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := make([][]int, 0)

	for i := 0; i < len(nums); i++ {

		target := -1 * nums[i]

		l := i + 1
		r := len(nums) - 1

		for l < r {
			num1 := nums[l]
			num2 := nums[r]
			sum := num1 + num2
			if sum == target {
				res = append(res, []int{nums[i], nums[l], nums[r]})

				for l+1 < r && nums[l+1] == nums[l] {
					l++
				}
				for r-1 > l && nums[r-1] == nums[r] {
					r--
				}
				l++
				r--
				continue

			}
			if sum < target {
				l++
			} else {
				r--
			}
		}
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}

	}

	return res

}
