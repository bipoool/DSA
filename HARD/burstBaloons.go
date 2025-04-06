package hard

var memCoin [][]int

func maxCoins(nums []int) int {

	l := 1
	r := len(nums)
	nums = append(nums, 1)
	finalNums := []int{1}
	finalNums = append(finalNums, nums...)

	memCoin = make([][]int, len(nums))
	for i := range memCoin {
		memCoin[i] = make([]int, len(nums))
		for j := range memCoin[i] {
			memCoin[i][j] = -1
		}
	}

	return dfsMaxCoins(finalNums, l, r)

}

func dfsMaxCoins(nums []int, l int, r int) int {

	if l > r {
		return 0
	}
	if memCoin[l][r] != -1 {
		return memCoin[l][r]
	}

	memCoin[l][r] = 0
	for i := l; i <= r; i++ {
		coins := nums[l-1] * nums[i] * nums[r+1]
		coins += dfsMaxCoins(nums, l, i-1) + dfsMaxCoins(nums, i+1, r)
		memCoin[l][r] = max(memCoin[l][r], coins)
	}
	return memCoin[l][r]

}
