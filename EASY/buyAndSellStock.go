package easy

import "math"

func maxProfit(prices []int) int {

	l := 0
	r := 0

	mxProfit := math.MinInt
	for r < len(prices) {

		profit := prices[r] - prices[l]
		if mxProfit < profit {
			mxProfit = profit
		}

		if prices[l] > prices[r] {
			l = r
			r++
		} else {
			r++
		}

	}

	return mxProfit
}

func main() {
	println(maxProfit(
		[]int{
			2, 1, 2, 1, 0, 1, 2,
		},
	))
}
