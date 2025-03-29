package medium

var resPer [][]int

func permute(nums []int) [][]int {

	resPer = make([][]int, 0)
	helperPer(nums, []int{}, 0, 0)
	return resPer

}

func helperPer(nums []int, currArr []int, curri int, currStart int) {

	if len(nums) == len(currArr) {
		temp := make([]int, len(currArr))
		copy(temp, currArr)
		resPer = append(resPer, temp)
		return
	}

	for i := currStart; i < len(nums); i++ {
		if i == curri {
			continue
		}
		currArr = append(currArr, nums[i])
		helperPer(nums, currArr, i, i)
		currArr = currArr[:len(currArr)-1]
	}

}
