package medium

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0
	res := l1

	for l1 != nil {

		sum := l1.Val + carry
		carry = 0

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum > 9 {
			carry = sum % 10
			sum = sum / 10
		}

		l1.Val = sum
		l1 = l1.Next

	}

	return res

}
