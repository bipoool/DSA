package easy

import "strings"

func isPalindrome(s string) bool {
	lowerString := ""
	for i := range s {
		lowerChar := strings.ToLower(string(s[i]))
		if lowerChar[0] >= 'a' && lowerChar[0] <= 'z' {
			lowerString = lowerString + lowerChar
		} else if lowerChar[0] >= '0' && lowerChar[0] <= '9' {
			lowerString = lowerString + lowerChar
		}
	}
	println(lowerString)
	println(lowerString)
	for i := range lowerString {
		if lowerString[i] != lowerString[len(lowerString)-i-1] {
			return false
		}
	}
	return true
}
