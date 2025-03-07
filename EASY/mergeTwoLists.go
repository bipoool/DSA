package easy

import DS "DSA/DS"

func mergeTwoLists(list1 *DS.ListNode, list2 *DS.ListNode) *DS.ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}

}

// func main() {
// 	l1 := DS.CreateListFromArr(
// 		[]int{
// 			1, 2, 4,
// 		},
// 	)
// 	l2 := DS.CreateListFromArr(
// 		[]int{
// 			1, 3, 4,
// 		},
// 	)

// 	DS.PrintList(l1)
// 	DS.PrintList(l2)

// 	res := mergeTwoLists(l1, l2)
// 	DS.PrintList(res)
// }
