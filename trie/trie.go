package trie

type TrieNode struct {
	next   map[rune]*TrieNode
	isWord bool
}

func NewTrieNode() *TrieNode {
	root := new(TrieNode)
	root.next = make(map[rune]*TrieNode)
	root.isWord = false
	return root
}

func (t *TrieNode) Insert(word string) {
	n := t
	for _, v := range word {
		if n.next[v] == nil {
			n.next[v] = NewTrieNode()
		}
		n = n.next[v]
	}
	n.isWord = true
}

func (t *TrieNode) Search(word string) bool {
	n := t
	for _, v := range word {
		if n.next[v] == nil {
			return false
		}
		n = n.next[v]
	}
	return n.isWord
}

func (t *TrieNode) StartWith(prefix string) bool {
	n := t
	for _, v := range prefix {
		if n.next[v] == nil {
			return false
		}
		n = n.next[v]
	}
	return true
}

