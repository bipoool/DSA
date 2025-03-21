package medium

func groupAnagrams(strs []string) [][]string {

	hmap := map[string][]string{}

	for i := range strs {
		str := strs[i]
		key := getKey(str)
		if _, ok := hmap[key]; !ok {
			hmap[key] = make([]string, 0)
		}
		hmap[key] = append(hmap[key], str)
	}

	res := make([][]string, 0)

	for i := range hmap {
		res = append(res, hmap[i])
	}

	return res

}

func getKey(str string) string {

	hmap := map[byte]int{}
	for i := range str {
		hmap[str[i]]++
	}

	key := "00000000000000000000000000"
	for i, v := range hmap {
		b := str[i] - 'a'
		key = key[:b] + string(v) + key[b+1:]
	}
	return key

}
