package medium

func checkValidString(s string) bool {

	leftMin := 0
	leftMax := 0

	for i := range s {
		char := s[i]

		if char == '(' {
			leftMax = leftMax + 1
			leftMin = leftMin + 1
		} else if char == ')' {
			leftMax = leftMax - 1
			leftMin = leftMin - 1
		} else {
			leftMax = leftMax + 1
			leftMin = leftMin - 1
		}
		if leftMax < 0 {
			return false
		}
		if leftMin < 0 {
			leftMin = 1
		}
	}
	return leftMin == 0

}
