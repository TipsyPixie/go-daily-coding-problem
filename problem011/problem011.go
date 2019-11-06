package problem011

import (
	"fmt"
	"strings"
)

type trie map[string]*trie

func NewTrie() (t *trie) {
	t = &trie{}
	return
}

func (node *trie) Insert(query string) *trie {
	cursor := node
	queryLeft := query
	for len(queryLeft) > 0 {
		if (*cursor)[queryLeft[:1]] == nil {
			(*cursor)[queryLeft[:1]] = NewTrie()
		}
		cursor = (*cursor)[queryLeft[:1]]
		queryLeft = queryLeft[1:]
	}
	return node
}

func (node *trie) Find(query string) []string {
	var prefixBuilder strings.Builder
	cursor := node
	queryLeft := query
	for len(queryLeft) > 0 && (*cursor)[queryLeft[:1]] != nil {
		prefixBuilder.WriteString(queryLeft[:1])
		cursor = (*cursor)[queryLeft[:1]]
		queryLeft = queryLeft[1:]
	}
	var result []string
	if len(queryLeft) == 0 {
		prefix := prefixBuilder.String()
		for _, substring := range cursor.subStrings() {
			result = append(result, fmt.Sprintf("%s%s", prefix, substring))
		}
	}
	return result
}

func (node *trie) subStrings() []string {
	if len(*node) == 0 {
		return []string{""}
	}
	childrenChannel := make(chan []string, len(*node))
	for keyString, subNode := range *node {
		go func(prefix string, node *trie) {
			var subStrings []string
			for _, subSubString := range node.subStrings() {
				subStrings = append(subStrings, fmt.Sprintf("%s%s", prefix, subSubString))
			}
			childrenChannel <- subStrings
		}(keyString, subNode)
	}
	finished := 0
	var result []string
	for finished < len(*node) {
		select {
		case subStrings := <-childrenChannel:
			result = append(result, subStrings...)
			finished++
		}
	}
	return result
}
