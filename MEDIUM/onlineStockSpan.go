package medium

// Store a monotonically increasing array with count of all the elements below it
type StockSpanner struct {
	st [][]int
}

func ConstructorStockSpanner() StockSpanner {
	return StockSpanner{
		st: [][]int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	count := 1
	for len(this.st) > 0 && this.st[len(this.st)-1][0] <= price {
		count += this.st[len(this.st)-1][1]
		this.st = this.st[:len(this.st)-1]
	}
	this.st = append(this.st, []int{price, count})
	return count
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
