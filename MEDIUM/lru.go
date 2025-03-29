package medium

type node struct {
	val  int
	key  int
	next *node
	prev *node
}

type LRUCache struct {
	hmap      map[int]*node
	nodeEnd   *node
	nodeStart *node
	len       int
	capacity  int
}

func ConstructorLRUCache(capacity int) LRUCache {
	lru := LRUCache{
		hmap:      map[int]*node{},
		nodeEnd:   &node{},
		nodeStart: &node{},
		len:       0,
		capacity:  capacity,
	}
	lru.nodeStart.next = lru.nodeEnd
	lru.nodeEnd.prev = lru.nodeStart
	return lru
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.hmap[key]; !ok {
		return -1
	}
	this.moveNodeToStart(key)
	return this.hmap[key].val
}

func (this *LRUCache) Put(key int, value int) {

	if _, ok := this.hmap[key]; !ok {
		this.hmap[key] = &node{}
		if this.capacity == this.len {
			this.evict()
			this.len--
		}
		this.len++
	}
	this.hmap[key].val = value
	this.hmap[key].key = key

	if this.hmap[key].prev != nil {
		this.hmap[key].prev.next = this.hmap[key].next
	}
	if this.hmap[key].next != nil {
		this.hmap[key].next.prev = this.hmap[key].prev
	}

	listPrev := this.nodeEnd.prev
	this.nodeEnd.prev = this.hmap[key]
	this.hmap[key].next = this.nodeEnd
	listPrev.next = this.hmap[key]
	this.hmap[key].prev = listPrev
}

func (this *LRUCache) moveNodeToStart(key int) {

	node := this.hmap[key]

	node.prev.next = node.next
	node.next.prev = node.prev

	listPrev := this.nodeEnd.prev
	this.nodeEnd.prev = node
	node.next = this.nodeEnd
	node.prev = listPrev
	listPrev.next = node

}

func (this *LRUCache) evict() {
	evictNode := this.nodeStart.next
	this.nodeStart.next = evictNode.next
	evictNode.prev = this.nodeStart
	delete(this.hmap, evictNode.key)
}
