package problem067

type freqNode struct {
	next       *freqNode
	prev       *freqNode
	refCount   int
	dNodesHead *dataNode
	dNodesTail *dataNode
}

type dataNode struct {
	next  *dataNode
	prev  *dataNode
	key   string
	value interface{}
	fNode *freqNode
}

type LfuCache struct {
	fNodesHead *freqNode
	fNodesTail *freqNode
	dNodes     map[string]*dataNode
	maxSize    uint
}

func NewCache(maxSize uint) *LfuCache {
	m := maxSize
	if m <= 0 {
		m = 1
	}
	h, t := &freqNode{
		refCount: -1,
	}, &freqNode{
		refCount: -1,
	}
	h.next, t.prev = t, h

	return &LfuCache{
		fNodesHead: h,
		fNodesTail: t,
		dNodes:     map[string]*dataNode{},
		maxSize:    m,
	}
}

func (c *LfuCache) Get(key string) interface{} {
	dNode, dNodeExists := c.dNodes[key]
	// return nil if the key not exists
	if !dNodeExists {
		return nil
	}

	fNode := dNode.fNode
	newFNode := fNode.next
	if newFNode.refCount != fNode.refCount+1 {
		// create a freq node with refCount = fNode.refCount + 1, if not any
		h, t := &dataNode{}, &dataNode{}
		h.next, t.prev = t, h

		newFNode = &freqNode{
			prev:       fNode,
			next:       fNode.next,
			refCount:   fNode.refCount + 1,
			dNodesHead: h,
			dNodesTail: t,
		}
		fNode.next, fNode.next.prev = newFNode, newFNode
	}

	dNode.prev.next, dNode.next.prev = dNode.next, dNode.prev
	newFNode.dNodesTail.prev, newFNode.dNodesTail.prev.next = dNode, dNode
	dNode.fNode, dNode.prev, dNode.next = newFNode, newFNode.dNodesTail.prev, newFNode.dNodesTail

	if fNode.dNodesHead.next == fNode.dNodesTail {
		fNode.prev.next, fNode.next.prev = fNode.next, fNode.prev
	}

	return dNode.value
}

func (c *LfuCache) Set(key string, value interface{}) {
	if _, alreadyExists := c.dNodes[key]; alreadyExists {
		return
	}

	for len(c.dNodes) >= int(c.maxSize) {
		// delete the least frequently used dNode
		lfuFNode, lfuDNode := c.fNodesHead.next, c.fNodesHead.next.dNodesHead.next
		delete(c.dNodes, lfuDNode.key)
		lfuDNode.prev.next, lfuDNode.next.prev = lfuDNode.next, lfuDNode.prev
		if lfuFNode.dNodesHead.next == lfuFNode.dNodesTail {
			lfuFNode.prev.next, lfuFNode.next.prev = lfuFNode.next, lfuFNode.prev
		}
	}

	fNode := c.fNodesHead.next
	if fNode.refCount != 0 {
		h, t := &dataNode{}, &dataNode{}
		h.next, t.prev = t, h

		fNode = &freqNode{
			next:       c.fNodesHead.next,
			prev:       c.fNodesHead,
			refCount:   0,
			dNodesHead: h,
			dNodesTail: t,
		}
		c.fNodesHead.next, c.fNodesHead.next.prev = fNode, fNode
	}

	dNode := &dataNode{
		next:  fNode.dNodesTail,
		prev:  fNode.dNodesTail.prev,
		key:   key,
		value: value,
		fNode: fNode,
	}
	fNode.dNodesTail.prev, fNode.dNodesTail.prev.next = dNode, dNode

	c.dNodes[key] = dNode
}
