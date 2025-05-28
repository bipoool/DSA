package medium

import "math"

// Create an array of length amount + 1 for memoization
// Loop in the amount and check how many coins would be required to make the current amount if we consider the current coin
// you can do this by checking mem[curAmount-coins[i]]
// mem[curAmount] = min(mem[curAmount] & mem[curAmount-coins[i]])
func coinChange(c []int, a int) int {

	ar := make([]int, a+1)
	for i := range ar {
		ar[i] = math.MaxInt
	}
	ar[0] = 0
	for j := 1; j <= a; j++ {
		for i := range c {
			if j-c[i] >= 0 && ar[j-c[i]] != math.MaxInt {
				ar[j] = min(ar[j], 1+ar[j-c[i]])
			}
		}
	}

	if ar[a] == math.MaxInt {
		return -1
	}
	return ar[a]
}

// DFS solution top down
var memcoinChange []int

func coinChangeDfs(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}
	memcoinChange = make([]int, amount+1)

	for i := range mem {
		mem[i] = -1
	}
	memcoinChange[0] = 0
	helpcoinChange(coins, amount)

	if memcoinChange[amount] == math.MaxInt {
		return -1
	}

	return memcoinChange[amount]

}

func helpcoinChange(coins []int, cursum int) int {

	if memcoinChange[cursum] != -1 {
		return memcoinChange[cursum]
	}

	res := math.MaxInt
	for i := range coins {
		if cursum-coins[i] >= 0 {
			r1 := helpcoinChange(coins, cursum-coins[i])
			if r1 == math.MaxInt {
				continue
			}
			res = min(res, 1+r1)
		}
	}

	memcoinChange[cursum] = res
	return res

}
