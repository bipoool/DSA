package easy

func twoSum(nums []int, target int) []int {
	var sumMap map[int]int = map[int]int{}
	for i := range nums {
		num2 := target - nums[i]
		if j, ok := sumMap[int(num2)]; ok {
			return []int{i, j}
		} else {
			sumMap[nums[i]] = i
		}
	}
	return nums
}
