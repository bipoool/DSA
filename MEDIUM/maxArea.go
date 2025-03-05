package medium

import "math"

func maxArea(height []int) int {

	l := 0
	r := len(height) - 1

	mxArea := math.MinInt

	for l < r {

		left := height[l]
		right := height[r]
		area := int(math.Min(float64(left), float64(right))) * (l - r)
		if area > mxArea {
			mxArea = area
		}
		if left > right {
			r--
		} else {
			l--
		}
	}

	return mxArea

}
