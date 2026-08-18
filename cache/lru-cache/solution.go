package main

func main() {

}

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

type LRUCache struct {
	capacity int
	nodes    map[int]*Node
	sentinel *Node
}

func Constructor(capacity int) LRUCache {
	sentinel := &Node{}
	sentinel.next = sentinel
	sentinel.prev = sentinel

	cache := LRUCache{
		capacity: capacity,
		nodes:    map[int]*Node{},
		sentinel: sentinel,
	}

	return cache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.nodes[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.nodes[key]; ok {
		this.moveToHead(node)
		node.value = value
		return
	}

	node := &Node{key: key, value: value}
	this.nodes[key] = node
	this.pushHead(node)
	if len(this.nodes) > this.capacity {
		tail := this.popTail()
		delete(this.nodes, tail.key)
	}
}

func (this *LRUCache) isEmpty() bool {
	return this.sentinel.next == this.sentinel
}

func (this *LRUCache) pushHead(n *Node) {
	n.next = this.sentinel
	n.prev = this.sentinel.prev
	this.sentinel.prev.next = n
	this.sentinel.prev = n
}

func (this *LRUCache) remove(n *Node) {
	n.next.prev = n.prev
	n.prev.next = n.next

	n.prev = nil
	n.next = nil
}

func (this *LRUCache) moveToHead(n *Node) {
	this.remove(n)
	this.pushHead(n)
}

func (this *LRUCache) popTail() *Node {
	if this.isEmpty() {
		return nil
	}

	oldest := this.sentinel.next
	this.remove(this.sentinel.next)
	return oldest
}
