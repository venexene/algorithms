package main

func main() {

}

type Node struct {
	key   int
	value int
	count int
	next  *Node
	prev  *Node
}

type LFUCache struct {
	capacity int
	nodes    map[int]*Node
	freqMap  map[int]*Node
	minFreq  int
}

func Constructor(capacity int) LFUCache {
	cache := LFUCache{
		capacity: capacity,
		nodes:    map[int]*Node{},
		freqMap:  map[int]*Node{},
		minFreq:  1,
	}

	cache.freqMap[cache.minFreq] = &Node{}
	cache.freqMap[cache.minFreq].next = cache.freqMap[cache.minFreq]
	cache.freqMap[cache.minFreq].prev = cache.freqMap[cache.minFreq]

	return cache
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.nodes[key]
	if !ok {
		return -1
	}
	node.count++
	this.promote(node, node.count, node.count-1)
	return node.value
}

func (this *LFUCache) Put(key int, value int) {
	if node, ok := this.nodes[key]; ok {
		node.value = value
		node.count++
		this.promote(node, node.count, node.count-1)
		return
	}

	node := &Node{key: key, value: value, count: 1}
	if len(this.nodes)+1 > this.capacity {
		tail := this.popTail()
		delete(this.nodes, tail.key)
	}
	this.nodes[key] = node
	this.minFreq = 1
	this.pushTail(node, this.minFreq)
}

func (this *LFUCache) promote(n *Node, new, old int) {
	if _, ok := this.freqMap[new]; !ok {
		this.freqMap[new] = &Node{}
		this.freqMap[new].next = this.freqMap[new]
		this.freqMap[new].prev = this.freqMap[new]
	}
	this.remove(n)
	if this.freqMap[old].next == this.freqMap[old] && old == this.minFreq {
		this.minFreq++
	}
	this.pushTail(n, new)
}

func (this *LFUCache) isEmpty() bool {
	return this.freqMap[this.minFreq].next == this.freqMap[this.minFreq]
}

func (this *LFUCache) pushTail(n *Node, freq int) {
	n.next = this.freqMap[freq]
	n.prev = this.freqMap[freq].prev
	this.freqMap[freq].prev.next = n
	this.freqMap[freq].prev = n
}

func (this *LFUCache) remove(n *Node) {
	n.next.prev = n.prev
	n.prev.next = n.next

	n.prev = nil
	n.next = nil
}

func (this *LFUCache) popTail() *Node {
	if this.isEmpty() {
		return nil
	}

	oldest := this.freqMap[this.minFreq].next
	this.remove(this.freqMap[this.minFreq].next)
	return oldest
}
