package medium

type NodeR struct {
	Val    int
	Next   *NodeR
	Random *NodeR
}

func copyRandomList(head *NodeR) *NodeR {

	hmap := map[int]*NodeR{}
	res := &NodeR{}
	currNodeR := &NodeR{}

	for head != nil {

		val := head.Val
		randomNodeR := head.Random

		if _, ok := hmap[randomNodeR.Val]; !ok {
			hmap[val] = &NodeR{
				Val: val,
			}
		}
		currNodeR.Next = hmap[val]
		if randomNodeR != nil {
			if _, ok := hmap[randomNodeR.Val]; !ok {
				hmap[randomNodeR.Val] = &NodeR{
					Val: randomNodeR.Val,
				}
			}
			currNodeR.Random = hmap[randomNodeR.Val]
		}
		if res == nil {
			res = currNodeR
		}
		head = head.Next
		currNodeR = currNodeR.Next
	}

	return res

}
