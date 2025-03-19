package hard

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	i := 0
	j := 0

	sortedArr := []int{}

	for i+j < len(nums1)+len(nums2) {
		if i < len(nums1) && j < len(nums2) && nums1[i] <= nums2[j] {
			sortedArr = append(sortedArr, nums1[i])
			i++
			continue
		} else if j < len(nums2) {
			sortedArr = append(sortedArr, nums2[j])
			j++
			continue
		} else if i < len(nums1) {
			sortedArr = append(sortedArr, nums1[i])
			i++
			continue
		}
		break
	}
	for i < len(nums1) {
		sortedArr = append(sortedArr, nums1[i])
		i++
	}
	for j < len(nums2) {
		sortedArr = append(sortedArr, nums2[j])
		j++
	}
	if len(sortedArr)%2 == 0 {
		mid := len(sortedArr) / 2
		sumArr := sortedArr[mid] + sortedArr[mid-1]
		return float64(sumArr) / 2
	}
	return float64(sortedArr[(len(sortedArr) / 2)])
}

// func main() {
// 	a := findMedianSortedArrays(
// 		[]int{4}, []int{1, 2, 3},
// 	)

// 	println(a)
// }
