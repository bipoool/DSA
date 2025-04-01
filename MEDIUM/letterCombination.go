package medium

var mapping map[byte][]string
var resLetter []string

func letterCombinations(digits string) []string {

	mapping['2'] = []string{"a", "b", "c"}
	mapping['3'] = []string{"d", "e", "f"}
	mapping['4'] = []string{"g", "h", "i"}
	mapping['5'] = []string{"j", "k", "l"}
	mapping['6'] = []string{"m", "n", "0"}
	mapping['7'] = []string{"p", "q", "r", "s"}
	mapping['8'] = []string{"t", "u", "v"}
	mapping['9'] = []string{"w", "x", "y", "x"}

	mapping = map[byte][]string{}
	resLetter = make([]string, 0)
	helperletterCombinations(digits, "", 0, 0)

	return resLetter

}

func helperletterCombinations(d string, currst string, st int, mst int) {

	if len(currst) == len(d) {
		resLetter = append(resLetter, currst)
		return
	}

	for i := st; i < len(d); i++ {
		for j := mst; j < len(mapping[d[i]]); j++ {
			helperletterCombinations(d, currst+mapping[d[i]][j], i+1, j+1)
		}
	}

}
