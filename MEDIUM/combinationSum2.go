package medium

import "sort"

var resCom [][]int

func combinationSum2(candidates []int, target int) [][]int {

	resCom = make([][]int, 0)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	helperCom2(candidates, target, []int{}, 0, 0)
	return res
}

func helperCom2(nums []int, target int, currArr []int, currSum int, currI int) {

	if currSum == target {
		temp := make([]int, len(currArr))
		copy(temp, currArr)
		resCom = append(resCom, temp)
	}
	if currSum > target {
		return
	}

	for i := currI; i < len(nums); i++ {
		currArr = append(currArr, nums[i])
		helperCom2(nums, target, currArr, currSum+nums[i], i+1)
		currArr = currArr[0 : len(currArr)-1]
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}

}
