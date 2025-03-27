package medium

func generateParenthesis(n int) []string {
	return helper([]string{}, "(", n, 1, 0)
}

func helper(res []string, currString string, n int, open int, close int) []string {
	if open == close && len(currString) == 2*n {
		res = append(res, currString)
		return res
	}
	if len(currString) > 2*n {
		return res
	}

	res = helper(res, currString+"(", n, open+1, close)

	if close < open {
		res = helper(res, currString+")", n, open, close+1)
	}

	return res
}
