package hard

import (
	"slices"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {

	if !slices.Contains(wordList, endWord) {
		return 0
	}
	hmapWord := map[string]int{}

	for i := range wordList {
		hmapWord[wordList[i]] = 1
	}

	que := []string{}
	qs := -1

	que = append(que, beginWord)
	qs++
	path := 0

	for qs < len(que) {
		path++
		qlen := len(que)
		for i := range qlen {
			word := que[i]
			for j := range que[i] {
				for k := range 26 {
					next := word[:j] + string(k+'a') + word[j+1:]
					if _, ok := hmapWord[next]; !ok {
						continue
					}
					delete(hmapWord, next)
					if next == endWord {
						return path
					}
				}
			}
		}
	}
	return 0
}
