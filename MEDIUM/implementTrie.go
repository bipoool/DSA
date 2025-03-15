package medium

type Trie struct {
	trieMap map[byte]*Trie
	isWord  bool
}

func ConstructorTrie() Trie {
	return Trie{
		trieMap: map[byte]*Trie{},
		isWord:  false,
	}
}

func (this *Trie) Insert(word string) {

	wordId := 0
	currTrie := this

	for wordId < len(word) {
		if trie, ok := currTrie.trieMap[word[wordId]]; ok {
			wordId++
			currTrie = trie
			continue
		}
		break
	}

	if wordId == len(word) {
		currTrie.isWord = true
		return
	}

	for i := wordId; i < len(word); i++ {
		letter := word[i]
		currTrie.trieMap[letter] = &Trie{isWord: false, trieMap: map[byte]*Trie{}}
		currTrie = currTrie.trieMap[letter]
	}
	currTrie.isWord = true
}

func (this *Trie) Search(word string) bool {
	wordId := 0
	currTrie := this
	for wordId < len(word) {
		if trie, ok := currTrie.trieMap[word[wordId]]; ok {
			wordId++
			currTrie = trie
			if wordId == len(word) {
				return currTrie.isWord
			}
		} else {
			break
		}
	}

	if wordId == len(word) {
		return currTrie.isWord
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	wordId := 0
	currTrie := this
	for wordId < len(prefix) {
		if trie, ok := currTrie.trieMap[prefix[wordId]]; ok {
			wordId++
			currTrie = trie
			if wordId == len(prefix) {
				return true
			}
		} else {
			break
		}
	}

	return false
}

// func main() {

// 	t := Constructor()
// 	t.Insert("hello")
// 	println(t.Search("hell"))
// 	println(t.Search("app"))
// 	t.Insert("app")
// 	println(t.Search("app"))

// }
