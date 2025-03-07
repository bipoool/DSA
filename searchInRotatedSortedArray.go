package main

func search(nums []int, target int) int {

	l := 0
	r := len(nums) - 1
	var mid int

	for l <= r {

		mid = (r + l) / 2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] < nums[0] {

			if nums[mid] < target && target <= nums[len(nums)-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {

			if nums[mid] > target && target >= nums[0] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

	}
	if nums[mid] != target {
		return -1
	}
	return nums[mid]
}

func main() {
	a := search([]int{
		4, 5, 6, 7, 0, 1, 2,
	}, 0)

	println(a)
}
