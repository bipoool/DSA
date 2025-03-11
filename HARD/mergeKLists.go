package hard

import "math"

func mergeKListsAppraoch2(lists []*ListNode) *ListNode {

	if lists == nil || len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return merge2List(lists[0], lists[1])
	}

	mid := len(lists) / 2
	l1 := mergeKListsAppraoch1(lists[0:mid])
	l2 := mergeKListsAppraoch2(lists[mid:])

	return merge2List(l1, l2)
}

func merge2List(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = merge2List(list1.Next, list2)
		return list1
	} else {
		list2.Next = merge2List(list1, list2.Next)
		return list2
	}

}

func mergeKListsAppraoch1(lists []*ListNode) *ListNode {

	if lists == nil || len(lists) == 0 {
		return nil
	}
	var index int
	minEle := math.MaxInt

	for i := range lists {

		if lists[i] != nil && lists[i].Val < minEle {
			minEle = lists[i].Val
			index = i
		}
	}

	if minEle == math.MaxInt {
		return nil
	}

	minList := lists[index]
	lists[index] = lists[index].Next
	minList.Next = mergeKListsAppraoch1(lists)

	return minList
}

type ListNode struct {
	Val  int
	Next *ListNode
}
