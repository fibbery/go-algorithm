package trie

import (
	"fmt"
	"testing"
)

func TestNewTrieNode(t *testing.T) {
	node := NewTrieNode()
	node.Insert("Hello")
	fmt.Println(node.Search("Hello"))
	fmt.Println(node.Search("Hallo"))
	fmt.Println(node.StartWith("Hel"))
}
