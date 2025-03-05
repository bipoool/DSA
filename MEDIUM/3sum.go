package medium

import "sort"

func threeSum(nums []int) [][]int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		target := -1 * nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left]+nums[right] == target {
				result = append(result, []int{nums[left], nums[right], nums[i]})
				for left < len(nums)-1 && nums[left] == nums[left+1] {
					left++
				}
				for right > 0 && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}
		}

		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return result

}
func main() {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	for _, v := range res {
		for _, j := range v {
			println(j)
		}
	}
}
