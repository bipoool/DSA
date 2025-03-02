package medium

func productExceptSelf(nums []int) []int {
	prod := 1
	foundZero := false
	foundZero2 := false
	for _, v := range nums {
		if v == 0 {
			if foundZero {
				foundZero2 = true
			} else {
				foundZero = true
			}
			continue
		}
		prod *= v
	}

	result := make([]int, 0)

	for _, v := range nums {
		if foundZero2 {
			result = append(result, 0)
		} else if v == 0 {
			result = append(result, prod)
		} else if foundZero {
			result = append(result, 0)
		} else {
			result = append(result, prod/v)
		}
	}

	return result

}
