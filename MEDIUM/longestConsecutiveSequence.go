package medium

func longestConsecutiveAppraoch1(nums []int) int {

	seqMap := map[int]int{}
	freqMap := map[int]int{}
	maxCount := map[int]int{}

	for _, v := range nums {
		freqMap[v] = 1
	}
	for _, num := range nums {
		seqMap = updatePrevAndForward(seqMap, freqMap, num, num)
	}

	maxCountNumber := 0
	for _, v := range seqMap {
		maxCount[v]++
		if maxCountNumber > maxCount[v] {
			maxCountNumber = maxCount[v]
		}
	}

	return maxCountNumber

}

func updatePrevAndForward(seqMap map[int]int, freqMap map[int]int, num int, group int) map[int]int {

	if _, ok := freqMap[num]; !ok {
		return seqMap
	}
	if _, ok := seqMap[num]; ok {
		return seqMap
	}

	seqMap[num] = group
	seqMap = updatePrevAndForward(seqMap, freqMap, num-1, group)
	seqMap = updatePrevAndForward(seqMap, freqMap, num+1, group)

	return seqMap
}

func longestConsecutiveAppraoch2(nums []int) int {

	freqMap := map[int]int{}

	for _, v := range nums {
		freqMap[v] = 1
	}
	maxCountNumber := 0
	for k := range freqMap {
		if _, ok := freqMap[k-1]; ok {
			continue
		}
		count := 1
		index := k
		for {
			if _, ok := freqMap[index+1]; ok {
				count++
				index++
			} else {
				break
			}
		}
		if maxCountNumber < count {
			maxCountNumber = count
		}
	}

	return maxCountNumber

}
