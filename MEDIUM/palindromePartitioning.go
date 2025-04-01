package medium

var resPart [][]string
var mem map[string]int

func partition(s string) [][]string {
	resPart = make([][]string, 0)
	mem = map[string]int{}
	helperPart(s, "", []string{}, 0)
	return resPart
}

func helperPart(s string, currStr string, currArr []string, start int) {

	if start == len(s) {
		temp := make([]string, len(currArr))
		copy(temp, currArr)
		resPart = append(resPart, temp)
	}

	for i := start; i < len(s); i++ {
		if isPalindrom(s[start : i+1]) {
			currArr = append(currArr, s[start:i+1])
			helperPart(s, s[start:i+1], currArr, i+1)
			currArr = currArr[:len(currArr)-1]
		}
	}

}

func isPalindrom(s string) bool {

	if mem[s] != 0 {
		return true
	}
	l := 0
	r := len(s) - 1

	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	mem[s] = 1
	return true

}
