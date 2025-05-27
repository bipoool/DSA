package medium

// First solve it without considering the last element
// Then solve it without considering the first element

// At each index you have 2 options
// Either rob the 2nd house
// Or rob what amount is possible if last house was not robbed
func rob(n []int) int {

	if len(n) == 1 {
		return n[0]
	}
	if len(n) == 2 {
		return max(n[0], n[1])
	}

	arr := make([]int, len(n))
	arr[0] = n[0]
	arr[1] = n[1]

	for i := 2; i < len(n)-1; i++ {
		arr[i] = max(arr[i-1]-n[i-1], arr[i-2]) + n[i]
	}

	res1 := max(arr[len(arr)-3], arr[len(arr)-2])
	arr[2] = n[2]
	arr[1] = n[1]
	for i := 3; i < len(n); i++ {
		arr[i] = max(arr[i-1]-n[i-1], arr[i-2]) + n[i]
	}

	res2 := max(arr[len(arr)-1], arr[len(arr)-2])
	return max(res1, res2)

}
