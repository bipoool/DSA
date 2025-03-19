package medium

func wordBreak(s string, wordDict []string) bool {
	dict := map[string]int{}
	dp := map[string]bool{}
	for i := range wordDict {
		dict[wordDict[i]] = 1
	}
	return helper(s, dict, dp)
}

func helper(s string, dict map[string]int, dp map[string]bool) bool {

	if len(s) == 0 {
		return true
	}
	if _, ok := dp[s]; ok {
		return dp[s]
	}
	word := ""
	for i := len(s) - 1; i >= 0; i-- {
		word = string(s[i:])
		if _, ok := dict[word]; ok {
			dp[word] = true
			if i-1 == 0 {
				return true
			}
			if helper(s[i+1:], dict, dp) {
				return true
			}
		}
	}
	return false

}
