package medium

type WordDictionary struct {
	trieMap map[byte]*WordDictionary
	isWord  bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		trieMap: map[byte]*WordDictionary{},
		isWord:  false,
	}
}

func (this *WordDictionary) AddWord(word string) {

	currTrie := this
	currIndex := 0

	for currIndex < len(word) {
		if trie, ok := currTrie.trieMap[word[currIndex]]; ok {
			currIndex++
			currTrie = trie
		} else {
			break
		}
	}
	for currIndex < len(word) {
		currTrie.trieMap[word[currIndex]] = &WordDictionary{
			trieMap: map[byte]*WordDictionary{},
		}
		currTrie = currTrie.trieMap[word[currIndex]]
		currIndex++
	}
	currTrie.isWord = true
}

func (this *WordDictionary) Search(word string) bool {

	currTrie := this
	currIndex := 0

	for currIndex = range word {
		if word[currIndex] == '.' {
			for j := range currTrie.trieMap {
				if currTrie.trieMap[j].Search(word[currIndex+1:]) {
					return true
				}
			}
			return false

		}

		if trie, ok := currTrie.trieMap[word[currIndex]]; ok {
			currIndex++
			currTrie = trie
		} else {
			break
		}
	}

	if currTrie.isWord && currIndex == len(word) {
		return true
	}
	return false
}

// func main() {
// 	dir := Constructor()
// 	dir.AddWord("a")
// 	dir.AddWord("a")
// 	println(dir.Search("a."))
// }
