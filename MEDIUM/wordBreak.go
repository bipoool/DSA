package medium

// Loop int string
// Check if we would have started from the current index, is there any word in dict that would have matched (loop in the dict for this)
// If it matches with on of the word, check if word after the current word also matches or not -> dp[i] = dp[i+len(w)]
func wordBreak(s string, wordDict []string) bool {

	dp := make([]int, len(s)+1)
	dp[len(s)] = 1

	for i := len(s) - 1; i >= 0; i-- {
		for _, w := range wordDict {
			if i+len(w) <= len(s) && s[i:i+len(w)] == w {
				dp[i] = dp[i+len(w)]
			}
			if dp[i] == 1 {
				break
			}
		}
	}
	return dp[0] == 1
}
