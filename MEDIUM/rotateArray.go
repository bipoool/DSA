package medium

func rotate(nums []int, k int) {
	k = k % len(nums)
	ex := nums[:k]
	st := nums[k:]
	rev(nums)
	rev(st)
	rev(ex)
}

func rev(n []int) {
	l := 0
	for l < len(n)/2 {
		n[l], n[len(n)-l-1] = n[len(n)-l-1], n[l]
		l++
	}
}
