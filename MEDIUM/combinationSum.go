package medium

func combinationSum(candidates []int, target int) [][]int {

	return helperCombinationSum(candidates, target, [][]int{}, []int{}, 0, 0)

}

func helperCombinationSum(arr []int, target int, resArr [][]int, currArr []int, currSum int, currIndex int) [][]int {

	if currSum == target {
		tempArr := make([]int, len(currArr))
		copy(tempArr, currArr)
		resArr = append(resArr, tempArr)
		return resArr
	}

	if currSum > target {
		return resArr
	}

	for i := currIndex; i < len(arr); i++ {
		currArr = append(currArr, arr[i])
		currSum += arr[i]
		resArr = helperCombinationSum(arr, target, resArr, currArr, currSum, i)
		currArr = currArr[:len(currArr)-1]
		currSum -= arr[i]
	}

	return resArr
}

// func main() {

// 	arr := []int{8, 7, 4, 3}
// 	res := combinationSum(arr, 11)

// 	for i := range res {

// 		for j := range res[i] {

// 			print(res[i][j])
// 			print(" ")

// 		}
// 		println()
// 	}

// }
