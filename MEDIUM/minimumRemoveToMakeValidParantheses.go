package medium

func minRemoveToMakeValid(s string) string {

	st := []int{}
	resArr := []rune(s)

	for i := range resArr {
		ch := s[i]
		if ch == '(' {
			st = append(st, len(resArr))
		} else if ch == ')' {
			if len(st) != 0 {
				st = st[:len(st)-1]
			} else {
				resArr[i] = 0
			}
		}
	}

	for len(st) != 0 {
		id := st[len(st)-1]
		st = st[:len(st)-1]
		resArr[id] = 0
	}

	resRune := make([]rune, 0)

	for i := range resArr {
		if resArr[i] != 0 {
			resRune = append(resRune, resArr[i])
		}
	}

	return string(resRune)

}
