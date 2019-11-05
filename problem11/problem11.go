package problem11

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
	for len(queryLeft) > 0 && (*cursor)[queryLeft[:1]] != nil {
		cursor = (*cursor)[queryLeft[:1]]
		queryLeft = queryLeft[1:]
	}
	for len(queryLeft) > 0 {
		(*cursor)[queryLeft[:1]] = NewTrie()
		cursor = (*cursor)[queryLeft[:1]]
		queryLeft = queryLeft[1:]
	}
	if (*cursor)[queryLeft] == nil {
		(*cursor)[queryLeft] = NewTrie()
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
		for _, substring := range cursor.getStrings() {
			result = append(result, fmt.Sprintf("%s%s", prefix, substring))
		}
	}
	return result
}

func (node *trie) getStrings() []string {
	if len(*node) == 0 {
		return []string{""}
	}
	childrenChannel := make(chan []string, len(*node))
	for key, value := range *node {
		go func(first string, subNode *trie) {
			var subStrings []string
			for _, subSubString := range subNode.getStrings() {
				subStrings = append(subStrings, fmt.Sprintf("%s%s", first, subSubString))
			}
			childrenChannel <- subStrings
		}(key, value)
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
