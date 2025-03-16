package medium

func rob1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	arr := make([]int, len(nums))
	arr[0] = nums[0]
	arr[1] = nums[1]

	for i := range nums {
		if i-2 < 0 {
			continue
		}
		if arr[i-2] < arr[i-1]-nums[i-1] {
			arr[i] = nums[i] + arr[i-1] - nums[i-1]
		} else {
			arr[i] = nums[i] + arr[i-2]
		}
	}
	if arr[len(nums)-1] > arr[len(nums)-2] {
		return arr[len(nums)-1]
	}
	return arr[len(nums)-2]
}
