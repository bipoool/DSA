package medium

import "math"

var dpProfit [][]int

func maxProfit(prices []int) int {

	dpProfit = make([][]int, len(prices))
	for i := range dpProfit {
		dpProfit[i] = make([]int, 2)
		dpProfit[i][0] = math.MinInt
		dpProfit[i][1] = math.MinInt
	}
	return dfsmaxProfit(prices, 0, 1)
}

func dfsmaxProfit(prices []int, i int, buy int) int {

	if i >= len(prices) {
		return 0
	}
	if dpProfit[i][buy] != math.MinInt {
		return dpProfit[i][buy]
	}

	coolDown := dfsmaxProfit(prices, i+1, buy)
	if buy == 1 {
		buyPr := dfsmaxProfit(prices, i+1, 0) - prices[i]
		dpProfit[i][buy] = max(coolDown, buyPr)
	} else {
		sellPr := dfsmaxProfit(prices, i+2, 1) + prices[i]
		dpProfit[i][buy] = max(coolDown, sellPr)
	}
	return dpProfit[i][buy]
}
