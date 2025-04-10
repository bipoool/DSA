package medium

func mergeTriplets(triplets [][]int, target []int) bool {

	res := []int{0, 0, 0}
	for i := range triplets {
		if triplets[i][0] > target[0] || triplets[i][1] > target[1] || triplets[i][2] > target[2] {
			continue
		}
		for j := range triplets[i] {
			if triplets[i][j] == target[j] {
				res[j] = 1
			}
		}
	}
	for i := range res {
		if res[i] != 1 {
			return false
		}
	}
	return true
}
