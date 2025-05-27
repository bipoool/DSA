package medium

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0
	res := l1
	prev := l1

	for l1 != nil {

		sum := l1.Val + carry
		carry = 0

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum > 9 {
			carry = sum / 10
			sum = sum % 10
		}

		l1.Val = sum
		prev = l1
		l1 = l1.Next
	}
	l1 = prev

	for l2 != nil {
		sum := l2.Val + carry
		carry = 0
		if sum > 9 {
			carry = sum / 10
			sum = sum % 10
		}
		l2.Val = sum
		prev = l1
		l1.Next = l2
		l2 = l2.Next
		l1 = l1.Next
	}

	if carry > 0 {
		l1.Next = &ListNode{
			Val: carry,
		}
	}

	return res

}
