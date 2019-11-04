package problem6

import (
	"errors"
	"math/rand"
	"time"
)

// Mocking pointer arithmetic
var mockMemory = make(map[NodePointer]*node)

type node struct {
	address NodePointer
	both    NodePointer
}

func (node node) getPointer() NodePointer {
	return node.address
}

type NodePointer uint

func (nodePointer NodePointer) dereferencePointer() *node {
	return mockMemory[nodePointer]
}

func NewNode(both NodePointer) *node {
	rand.Seed(time.Now().UnixNano())
	node := node{
		address: NodePointer(rand.Intn(1000000)) + 1,
		both:    both,
	}
	mockMemory[node.address] = &node
	return &node
}

func Add(head *node) error {
	if head == nil {
		return errors.New("cannot dereference nil head")
	}
	prevNodePointer, atCursor := NodePointer(0), head
	for atCursor.both != prevNodePointer {
		prevNodePointer, atCursor = atCursor.getPointer(), (prevNodePointer ^ atCursor.both).dereferencePointer()
	}
	newLastNode := NewNode(atCursor.getPointer())
	atCursor.both = prevNodePointer ^ newLastNode.getPointer()
	return nil
}

type IndexOutOfRange struct{}

func (IndexOutOfRange) Error() string {
	return "index out of range"
}

func Get(head *node, index int) (*node, error) {
	if head == nil || index < 0 {
		return nil, IndexOutOfRange{}
	}
	prevNodePointer, atCursor := NodePointer(0), head
	for i := 0; i < index; i++ {
		if atCursor.both == prevNodePointer {
			return nil, IndexOutOfRange{}
		}
		prevNodePointer, atCursor = atCursor.getPointer(), (prevNodePointer ^ atCursor.both).dereferencePointer()
	}
	return atCursor, nil
}
