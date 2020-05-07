package trie

import (
	"testing"
)

func TestTrie_Search(t *testing.T) {
	root := Constructor()
	root.Insert("apple")
	if got := root.Search("apple"); got != true{
		t.Errorf("Search() = %v, want %v", got, true)
	}
}

func TestTrie_StartsWith(t *testing.T) {
	root := Constructor()
	root.Insert("apple")
	if got := root.StartsWith("app"); got != true{
		t.Errorf("Search() = %v, want %v", got, false)
	}
}
