package easy

// Change buy price daily if currprice is less than old buy price
// Sell daily and save the maximum of mxProfit and old profit
func maxProfit(prices []int) int {
	mx := 0
	b := prices[0]

	for i := 1; i < len(prices); i++ {
		p := prices[i]
		curPr := p - b
		mx = max(mx, curPr)
		if b > p {
			b = p
		}
	}

	return mx

}
func main() {
	println(maxProfit(
		[]int{
			2, 1, 2, 1, 0, 1, 2,
		},
	))
}
