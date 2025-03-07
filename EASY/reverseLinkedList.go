package easy

import DS "DSA/DS"

func reverse(head *DS.ListNode, currNode *DS.ListNode) (*DS.ListNode, *DS.ListNode) {
	if currNode.Next == nil {
		return currNode, currNode
	}
	resHead, node := reverse(currNode.Next, currNode.Next)
	node.Next = head
	head.Next = nil
	return resHead, head
}
func reverseList(head *DS.ListNode) *DS.ListNode {
	res, _ := reverse(head, head)
	return res
}

// func main() {
// 	list := createListFromArr(
// 		[]int{
// 			1, 2, 3, 4, 5,
// 		},
// 	)
// 	printList(list)
// 	a := reverseList(list)
// 	printList(a)
// }
