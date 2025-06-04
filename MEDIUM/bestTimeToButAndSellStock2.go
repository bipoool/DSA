package medium

// Just loop from index 1
// Add profit -> cur-prev if prev < cur
func maxProfit2(prices []int) int {
	runpr := 0

	for i := 1; i < len(prices); i++ {
		prv := prices[i-1]
		cur := prices[i]

		if prv < cur {
			runpr += cur - prv
		}
	}

	return runpr

}
