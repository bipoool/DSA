package medium

var res [][]int

func subsets(nums []int) [][]int {
	res = make([][]int, 0)
	helperSub(nums, []int{}, 0)
	return res
}

func helperSub(nums []int, currArr []int, currI int) {
	temp := make([]int, len(currArr))
	copy(temp, currArr)
	res = append(res, temp)
	for i := currI; i < len(nums); i++ {
		currArr = append(currArr, nums[i])
		helperSub(nums, currArr, i+1)
		currArr = currArr[:len(currArr)-1]
	}
}
