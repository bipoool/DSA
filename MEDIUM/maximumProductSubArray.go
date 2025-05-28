package medium

import "math"

// Just calculate the prod from start and then from end
// Compare maxProd at each step
// If currProd == 0 -> currProd=1
func maxProduct(nums []int) int {
	mxProd := math.MinInt

	prod := 1

	for i := range nums {
		prod *= nums[i]
		mxProd = max(mxProd, prod)
		if nums[i] == 0 {
			prod = 1
		}
	}
	prod = 1
	for i := len(nums) - 1; i >= 0; i-- {
		prod *= nums[i]
		mxProd = max(mxProd, prod)
		if nums[i] == 0 {
			prod = 1
		}
	}
	return mxProd
}

// func main() {
// 	println(maxProduct(
// 		[]int{2, -5, -2, -4, 3},
// 	))
// }
