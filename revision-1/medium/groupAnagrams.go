package medium

func groupAnagrams(strs []string) [][]string {

	hmap := map[string][]string{}
	for _, str := range strs {
		key := getKey(str)
		if _, ok := hmap[key]; !ok {
			hmap[key] = []string{}
		}
		hmap[key] = append(hmap[key], str)
	}

	res := [][]string{}

	for _, v := range hmap {
		res = append(res, v)
	}
	return res

}

func getKey(str string) string {
	key := ""
	keyArr := [26]int{}
	for i := range str {
		keyArr[str[i]-'a']++
	}
	for k, v := range keyArr {
		key = key + string(v) + string(k+'a')
	}
	return key
}
