package medium

func canJump(nums []int) bool {

	mem := make([]bool, len(nums))
	mem[len(nums)-1] = true

	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j <= nums[i]; j++ {
			if mem[i] {
				break
			}
			step := i + j
			if step < len(nums) && mem[step] {
				mem[i] = true
			}
		}
	}

	return mem[0]

}
