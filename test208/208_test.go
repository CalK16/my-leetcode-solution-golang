package test208

import "testing"

type Trie struct {
	IsEnd    bool
	Val      rune
	Children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{Children: make(map[rune]*Trie)}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, c := range word {
		if curr.Children[c] == nil {
			curr.Children[c] = &Trie{Val: c, Children: make(map[rune]*Trie)}
		}
		curr = curr.Children[c]
	}
	curr.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, c := range word {
		if child := curr.Children[c]; child == nil {
			return false
		}
		curr = curr.Children[c]
	}
	return curr.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, c := range prefix {
		if child := curr.Children[c]; child == nil {
			return false
		}
		curr = curr.Children[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
func Test208_1(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("TEST FAILED")
	} // return True

	if trie.Search("app") {
		t.Errorf("TEST FAILED")
	} // return False

	if !trie.StartsWith("app") {
		t.Errorf("TEST FAILED")
	} // return True

	trie.Insert("app")

	if !trie.Search("app") {
		t.Errorf("TEST FAILED")
	} // return True
}
