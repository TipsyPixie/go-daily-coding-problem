package problem011

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("dog").Insert("dear").Insert("deal").Insert("demon").Insert("dope")
	resultMap := map[string]bool{}
	for _, result := range trie.Find("de") {
		resultMap[result] = true
	}
	if !resultMap["dear"] || !resultMap["deal"] || !resultMap["demon"] {
		t.Log(resultMap)
		t.FailNow()
	}
}
