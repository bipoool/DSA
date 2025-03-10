package medium

func reorderList(head *ListNode) {
	reorderHelper(head, head)
}

func reorderHelper(head *ListNode, nextNode *ListNode) *ListNode {
	if nextNode == nil {
		return head
	}
	headNode := reorderHelper(head, nextNode.Next)

	if headNode == nil || headNode == nextNode || nextNode.Next == headNode {
		if headNode != nil {
			headNode.Next = nil
		}
		return nil
	}

	headNext := headNode.Next
	headNode.Next = nextNode
	nextNode.Next = headNext

	return headNext
}

type ListNode struct {
	Val  int
	Next *ListNode
}
