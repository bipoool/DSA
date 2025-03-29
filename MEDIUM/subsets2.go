package medium

import "sort"

var ressub2 [][]int

func subsetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ressub2 = make([][]int, 0)
	helperSubSet2(nums, []int{}, 0)
	return ressub2
}

func helperSubSet2(nums []int, currArr []int, currI int) {

	if len(currArr) > len(nums) {
		return
	}
	temp := make([]int, len(currArr))
	copy(temp, currArr)
	ressub2 = append(ressub2, temp)

	for i := currI; i < len(nums); i++ {
		currArr = append(currArr, nums[i])
		helperSubSet2(nums, currArr, i+1)
		currArr = currArr[:len(currArr)-1]
		for i+1 < len(nums) && nums[i+1] == nums[i] {
			i++
		}
	}

}
