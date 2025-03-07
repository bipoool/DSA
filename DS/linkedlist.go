package ds

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(head *ListNode) {
	for head != nil {
		print(head.Val)
		print("->")
		head = head.Next
	}
	println("nil")
}

func CreateListFromArr(list []int) *ListNode {
	head := &ListNode{}
	res := head
	for _, v := range list {
		head.Next = &ListNode{}
		head.Next.Val = v
		head = head.Next
	}
	return res.Next
}
