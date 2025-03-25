package hard

import "math"

func largestRectangleArea(heights []int) int {

	maxstack := make([][]int, len(heights)+1)
	maxs := -1
	maxArea := math.MinInt

	for i := 0; i < len(heights); i++ {
		lastIn := i
		for maxs >= 0 && maxstack[maxs][1] >= heights[i] {
			height := maxstack[maxs][1]
			base := i - maxstack[maxs][0]
			maxArea = max(maxArea, base*height)
			lastIn = maxstack[maxs][0]
			maxs--
		}
		maxstack[maxs+1] = []int{lastIn, heights[i]}
		maxs++
	}

	for maxs >= 0 {
		base := len(heights) - maxstack[maxs][0]
		height := maxstack[maxs][1]
		maxArea = max(maxArea, base*height)
		maxs--
	}

	return maxArea

}
