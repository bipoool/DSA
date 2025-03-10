package medium

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	temp := &ListNode{
		Val:  0,
		Next: head,
	}
	dummy := temp

	for n > 0 {
		head = head.Next
	}

	for head != nil {
		head = head.Next
		dummy = dummy.Next
	}

	dummy.Next = dummy.Next.Next

	return temp.Next

}
