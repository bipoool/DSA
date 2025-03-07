package easy

func isValid(s string) bool {

	stack := make([]byte, len(s))
	i := 0

	for c := range s {
		if isOpening(s[c]) {
			stack[i] = s[c]
			i++
		} else if i > 0 {
			open := stack[i-1]
			i--
			if open == '{' && s[c] != byte('}') {
				return false
			}
			if open == '(' && s[c] != byte(')') {
				return false
			}
			if open == '[' && s[c] != byte(']') {
				return false
			}

		} else {
			return false
		}
	}

	return i == 0
}

func isOpening(c byte) bool {
	if c == byte('}') || c == byte(']') || c == byte(')') {
		return false
	}
	return true
}
