package no146

type LRUCache struct {
	capacity   int
	size       int
	values     map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		values:   make(map[int]*DLinkedNode),
		head:     &DLinkedNode{key: 0, value: 0},
		tail:     &DLinkedNode{key: 0, value: 0},
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.values[key]
	if ok {
		this.moveToHead(value)
		return value.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.values[key]; !ok {
		node := &DLinkedNode{key: key, value: value}
		this.values[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.values, removed.key)
			this.size--
		}
	} else {
		node := this.values[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
