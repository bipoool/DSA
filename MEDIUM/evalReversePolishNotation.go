package medium

import "strconv"

func evalRPN(tokens []string) int {

	stack := make([]int, len(tokens))
	start := -1

	for i := 0; i < len(tokens); i++ {
		if isDig(tokens[i]) {
			val, _ := strconv.ParseInt(tokens[i], 10, 0)
			stack[start+1] = int(val)
		} else {
			dig1 := stack[start]
			dig2 := stack[start-1]
			start -= 2
			res := eval(dig1, dig2, tokens[i])
			stack[start+1] = res
			start++
		}
	}
	return stack[start]

}

func eval(a int, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return a / b
	}
}

func isDig(a string) bool {

	switch a {
	case "+":
		return false
	case "-":
		return false
	case "*":
		return false
	case "/":
		return false
	default:
		return true
	}

}
