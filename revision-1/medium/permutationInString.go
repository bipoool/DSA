package medium

func checkInclusion(s1 string, s2 string) bool {

	hmap1 := map[byte]int{}
	hmap2 := map[byte]int{}

	for i := range s1 {
		hmap1[s1[i]] += 1
		hmap2[s2[i]] += 1
	}

	l := 0
	r := len(s1)
	for l < r {
		if compareMaps(hmap1, hmap2) {
			return true
		}

		r++
		if r < len(s2) {
			hmap2[s2[r]]++
		} else {
			break
		}
		hmap2[s2[l]]--
		l++
	}
	return false
}

func compareMaps(hmap1 map[byte]int, hmap2 map[byte]int) bool {
	for i := range hmap2 {
		if _, ok := hmap1[i]; !ok || hmap1[i] != hmap2[i] {
			return false
		}
	}
	return true
}

func main() {
	checkInclusion("abc", "bbbca")
}
