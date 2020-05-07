package trie

type Trie struct {
	value    string
	children map[string]*Trie
	isWord   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		value:    "",
		children: make(map[string]*Trie),
		isWord:   false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, value := range word {
		if _, ok := node.children[string(value)]; ok {
			// traverse next node while current node is exist
			node = node.children[string(value)]
		} else {
			newNode := Trie{
				value:    string(value),
				children: make(map[string]*Trie),
				isWord:   false,
			}
			node.children[string(value)] = &newNode
			node = node.children[string(value)]
		}
	}
	node.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, value := range word {
		if _, ok := node.children[string(value)]; ok {
			node = node.children[string(value)]
		}else {
			return false
		}
	}
	return node.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, value := range prefix {
		if _, ok := node.children[string(value)]; ok {
			node = node.children[string(value)]
		}else {
			return false
		}
	}
	return node.isWord || len(node.children) > 0
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
