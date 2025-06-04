package medium

import "github.com/emirpasic/gods/queues/priorityqueue"

// Create 2 heaps to store max and min element
// When geting the max and min element
// Check if the value matches from map or not
// If not pop it
// Else return the pop value
type Price struct {
	t int
	p int
}

type StockPrice struct {
	m  map[int]Price
	mx *priorityqueue.Queue
	mn *priorityqueue.Queue
	lt int
}

func ConstructorStockPrice() StockPrice {
	return StockPrice{
		m: map[int]Price{},
		mx: priorityqueue.NewWith(func(a, b interface{}) int {
			return b.(Price).p - a.(Price).p
		}),
		mn: priorityqueue.NewWith(func(a, b interface{}) int {
			return a.(Price).p - b.(Price).p
		}),
	}
}

func NewPrice(t int, p int) Price {
	return Price{
		t: t,
		p: p,
	}
}

func (this *StockPrice) Update(t int, p int) {

	price := NewPrice(t, p)
	this.mx.Enqueue(price)
	this.mn.Enqueue(price)
	this.lt = max(this.lt, t)
	this.m[t] = price

}

func (this *StockPrice) Current() int {
	return this.m[this.lt].p
}

func (this *StockPrice) Maximum() int {
	for this.mx.Size() > 0 {
		priceRaw, _ := this.mx.Peek()
		price := priceRaw.(Price)
		if price.p != this.m[price.t].p {
			this.mx.Dequeue()
		} else {
			break
		}
	}
	priceRaw, _ := this.mx.Peek()
	return priceRaw.(Price).p
}

func (this *StockPrice) Minimum() int {
	for this.mn.Size() > 0 {
		priceRaw, _ := this.mn.Peek()
		price := priceRaw.(Price)
		if price.p != this.m[price.t].p {
			this.mn.Dequeue()
		} else {
			break
		}
	}
	priceRaw, _ := this.mn.Peek()
	return priceRaw.(Price).p
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
