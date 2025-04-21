package medium

import "math"

// use Binary Seach
// l = 1 & r = max(piles)
// check if coco can finish all the bananas in mid speed
// Don't forget to use math.Ceil here
func minEatingSpeed(piles []int, h int) int {

	maxPile := math.MinInt

	for i := range piles {
		maxPile = max(maxPile, piles[i])
	}

	l := 0
	r := maxPile
	res := 0

	for l < maxPile {

		es := (r + l) / 2
		currSum := 0

		for i := range piles {
			currSum += int(math.Ceil(float64(piles[i]) / float64(es)))
		}

		if currSum <= h {
			res = currSum
			r = es - 1
		} else {
			l = l + 1
		}
	}
	return res
}
