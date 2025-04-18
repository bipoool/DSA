package medium

// Create map for both strings
// Compare the maps for each window of size len(s1)
// Comparing is only o(26)
func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}
	h1 := map[byte]int{}
	h2 := map[byte]int{}

	for i := range s1 {
		h1[s1[i]]++
		h2[s2[i]]++
	}

	l := 0

	for r := len(s1); r < len(s2); r++ {
		if compareMap(h1, h2) {
			return true
		}
		h2[s2[r]]++
		h2[s2[l]]--
		l++
	}

	return compareMap(h1, h2)

}

func compareMap(h1, h2 map[byte]int) bool {

	for k, v := range h1 {
		if h2[k] != v {
			return false
		}
	}
	return true

}
