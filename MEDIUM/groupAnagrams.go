package medium

func groupAnagrams(strs []string) [][]string {

	mapOfAnagrams := make(map[string][]string)
	result := make([][]string, 0)
	for _, str := range strs {
		key := findKey(str)
		if _, ok := mapOfAnagrams[key]; !ok {
			mapOfAnagrams[key] = make([]string, 0)
		}
		mapOfAnagrams[key] = append(mapOfAnagrams[key], str)
	}
	for _, val := range mapOfAnagrams {
		result = append(result, val)
	}
	return result
}

func findKey(str string) string {
	key := []rune("########################")
	keyMap := make(map[int]int)
	for i := range str {
		keyMap[(int(str[i]))-97] = keyMap[(int(str[i]))-97] + 1
	}

	for k, v := range keyMap {
		key[k] = rune(v)
	}

	return string(key)
}
