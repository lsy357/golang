package lru

import "container/list"

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}
type Pair struct {
	key   int
	value int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (lru *LRUCache) Get(key int) int {
	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem)
		elem.Value = Pair{key, value}
	} else {
		if lru.list.Len() >= lru.capacity {
			delete(lru.cache, lru.list.Back().Value.(Pair).key)
			lru.list.Remove(lru.list.Back())
		}
		lru.list.PushFront(Pair{key, value})
		lru.cache[key] = lru.list.Front()
	}
}
