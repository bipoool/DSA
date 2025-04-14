package medium

func productExceptSelf(nums []int) []int {

	zero := 0
	prod := 1
	for _, i := range nums {

		if i == 0 {
			zero++
			continue
		}

		prod *= i

	}

	if zero > 1 {
		return make([]int, len(nums))
	}

	res := make([]int, len(nums))

	for i := range nums {
		if zero > 0 {
			if nums[i] == 0 {
				res = append(res, prod)
			}
		} else {
			res = append(res, prod/nums[i])
		}

	}
	return res
}
