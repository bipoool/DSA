package main

import "math"

func mergeKLists(lists []*ListNode) *ListNode {

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
	minList.Next = mergeKLists(lists)

	return minList
}

type ListNode struct {
	Val  int
	Next *ListNode
}
