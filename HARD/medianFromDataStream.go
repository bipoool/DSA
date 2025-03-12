package hard

import (
	"container/heap"
)

type MinPq []int
type MaxPq []int

func (p *MinPq) Push(e any) {
	*p = append(*p, e.(int))
}

func (p *MinPq) Pop() any {
	old := *p
	res := old[len(*p)-1]
	*p = old[0 : len(*p)-1]
	return res
}

func (p *MinPq) Peak() any {
	old := *p
	return old[0]
}

func (p MinPq) Len() int      { return len(p) }
func (h MinPq) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinPq) Less(i, j int) bool {
	return h[i] < h[j]
}

func (p *MaxPq) Push(e any) {
	*p = append(*p, e.(int))
}

func (p *MaxPq) Pop() any {
	old := *p
	res := old[len(*p)-1]
	*p = old[0 : len(*p)-1]
	return res
}

func (p *MaxPq) Peak() any {
	old := *p
	return old[0]
}

func (p MaxPq) Len() int      { return len(p) }
func (h MaxPq) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MaxPq) Less(i, j int) bool {
	return h[i] > h[j]
}

type MedianFinder struct {
	pqRight *MinPq
	pqLeft  *MaxPq
}

func GetMedianFinder() MedianFinder {

	pqLeft := &MaxPq{}
	heap.Init(pqLeft)

	pqRight := &MinPq{}
	heap.Init(pqRight)

	return MedianFinder{
		pqLeft:  pqLeft,
		pqRight: pqRight,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.pqLeft.Len() == 0 {
		heap.Push(this.pqLeft, num)
		return
	} else if this.pqRight.Len() == 0 && this.pqLeft.Peak().(int) > num {
		leftEle := heap.Pop(this.pqLeft)
		heap.Push(this.pqRight, leftEle)
		heap.Push(this.pqLeft, num)
		return
	} else if this.pqRight.Len() == 0 {
		heap.Push(this.pqRight, num)
		return
	}
	rightEle := this.pqRight.Peak().(int)

	if rightEle <= num {
		heap.Push(this.pqRight, num)
	} else {
		heap.Push(this.pqLeft, num)
	}

	if this.pqRight.Len()-this.pqLeft.Len() > 1 {
		ele := heap.Pop(this.pqRight)
		heap.Push(this.pqLeft, ele)
	} else if this.pqLeft.Len()-this.pqRight.Len() > 1 {
		ele := heap.Pop(this.pqLeft)
		heap.Push(this.pqRight, ele)
	}

}

func (this *MedianFinder) FindMedian() float64 {

	if this.pqLeft.Len() == this.pqRight.Len() {
		return (float64(this.pqLeft.Peak().(int)) + float64(this.pqRight.Peak().(int))) / 2.0
	}

	if this.pqLeft.Len() > this.pqRight.Len() {
		return float64(this.pqLeft.Peak().(int))
	}

	return float64(this.pqRight.Peak().(int))
}

// func main() {
// 	m := GetMedianFinder()
// 	commands := []string{"addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"}
// 	arguments := []int{40, 0, 12, 0, 16, 0, 14, 0, 35, 0, 19, 0, 34, 0, 35, 0, 28, 0, 35, 0, 26, 0, 6, 0, 8, 0, 2, 0, 14, 0, 25, 0, 25, 0, 4, 0, 33, 0, 18, 0, 10, 0, 14, 0, 27, 0, 3, 0, 35, 0, 13, 0, 24, 0, 27, 0, 14, 0, 5, 0, 0, 0, 38, 0, 19, 0, 25, 0, 11, 0, 14, 0, 31, 0, 30, 0, 11, 0, 31, 0, 0, 0}

// 	for i, cmd := range commands {
// 		if cmd == "addNum" {
// 			m.AddNum(arguments[i])
// 		} else {
// 			println(m.FindMedian())
// 		}
// 	}

// }
