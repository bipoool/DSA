package hard

func reverseKGroup(head *ListNode, k int) *ListNode {
	start := head
	var prev *ListNode
	end := head
	var resHead *ListNode
	i := 0

	var lastGroupEnd *ListNode

	for end != nil {
		i++
		temp := end.Next
		if i%k == 0 {
			prev = reverse(start, prev, temp)
			if resHead == nil {
				resHead = end
			}
			if lastGroupEnd != nil {
				lastGroupEnd.Next = end
			}
			lastGroupEnd = prev
			start = temp
		}
		end = temp
	}
	lastGroupEnd.Next = start

	return resHead
}

func reverse(start *ListNode, prev *ListNode, end *ListNode) *ListNode {
	head := start
	for start != end {
		temp := start.Next
		start.Next = prev
		prev = start
		start = temp
	}
	return head
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

func main() {

	head := CreateListFromArr([]int{1, 2, 3, 4, 5})
	reverseKGroup(head, 2)

}
