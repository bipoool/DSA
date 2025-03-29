package hard

// Monotonically Decreasing Queue
// Only add elements which are smaller than the element in queue
// We will always have max element at the start of queue

func maxSlidingWindow(nums []int, k int) []int {

	queue := make([]int, len(nums))
	start := 0
	end := 0

	res := make([]int, 0)

	for i := range nums {

		for end > 0 && end > start && nums[i] > queue[end-1] {
			end--
		}
		queue[end] = nums[i]
		end++
		if i >= k && nums[i-k] == queue[start] {
			start = min(end-1, start+1)
		}
		if i >= k-1 {
			res = append(res, queue[start])
		}
	}
	return res
}

// func main() {
// 	maxSlidingWindow(
// 		[]int{9, 10, 9, -7, -4, -8, 2, -6}, 5,
// 	)
// }
