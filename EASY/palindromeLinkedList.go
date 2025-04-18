package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindromeLinkedList(head *ListNode) bool {

	half, h := getHalf(head)

	if h == nil || half == nil {
		return true
	}

	for h != nil && half != nil {
		if h.Val != half.Val {
			return false
		}
		h = h.Next
		half = half.Next
	}
	return h == nil && half == nil
}

func getHalf(h *ListNode) (*ListNode, *ListNode) {
	count := 0
	head := h
	for head != nil {
		count++
		head = head.Next
	}

	length := count
	end := count / 2
	count = 0

	var newHead *ListNode
	var prev *ListNode
	head = h

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next

		if count+1 == end {
			newHead = next
			break
		}
		count++
	}

	if length%2 == 0 || length == 1 {
		return newHead, prev
	}
	return newHead.Next, prev
}
