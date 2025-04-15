package easy

func merge(nums1 []int, m int, nums2 []int, n int) {

	l1 := m - 1
	l2 := n - 1
	rp := m + n - 1

	for l1 >= 0 && l2 >= 0 {
		if nums1[l1] > nums2[l2] {
			nums1[rp] = nums1[l1]
			l1--
		} else {
			nums1[rp] = nums2[l2]
			l2--
		}
		rp--
	}

	for l2 >= 0 {
		nums1[rp] = nums2[l2]
		l2--
		rp--
	}

}
