package problem052

type cacheKey string

type cacheValue string

// using mapped doubly linked list to satisfy O(1) complexity

type node struct {
	prev  *node
	next  *node
	key   cacheKey
	value *cacheValue
}

type cache struct {
	capacity  int
	length    int
	data      map[cacheKey]*node
	firstData *node
	lastData  *node
}

func NewCache(capacity int) *cache {
	return &cache{
		capacity:  capacity,
		length:    0,
		data:      map[cacheKey]*node{},
		firstData: nil,
		lastData:  nil,
	}
}

func (thisCache *cache) Get(key cacheKey) *cacheValue {
	if valueNode, valueExists := thisCache.data[key]; !valueExists {
		return nil
	} else {
		if thisCache.lastData == valueNode {
			return valueNode.value
		}
		if valueNode.prev != nil {
			valueNode.prev.next = valueNode.next
		}
		if valueNode.next != nil {
			valueNode.next.prev = valueNode.prev
		}
		thisCache.lastData.next = valueNode
		thisCache.lastData = valueNode
		return valueNode.value
	}
}

func (thisCache *cache) Set(key cacheKey, value cacheValue) *cache {
	if valueNode, valueExists := thisCache.data[key]; valueExists {
		valueNode.value = &value
	} else {
		if thisCache.firstData != nil && thisCache.length >= thisCache.capacity {
			if thisCache.firstData.next != nil {
				thisCache.firstData.next.prev = nil
			}
			delete(thisCache.data, thisCache.firstData.key)
			thisCache.firstData = thisCache.firstData.next
			thisCache.length--
		}

		valueNode = &node{
			prev:  thisCache.lastData,
			next:  nil,
			key:   key,
			value: &value,
		}
		if thisCache.lastData != nil {
			thisCache.lastData.next = valueNode
		}
		thisCache.lastData = valueNode
		thisCache.data[key] = valueNode
		thisCache.length++

		if thisCache.firstData == nil {
			thisCache.firstData = valueNode
		}
	}
	return thisCache
}
