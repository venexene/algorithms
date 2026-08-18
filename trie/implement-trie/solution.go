package main

func main() {

}

type Trie struct {
	root *Node
}

type Node struct {
	children [26]*Node
	isEnd    bool
}

func Constructor() Trie {
	return Trie{
		root: &Node{},
	}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, r := range word {
		if node.children[r-'a'] == nil {
			node.children[r-'a'] = &Node{}
		}
		node = node.children[r-'a']
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, r := range word {
		node = node.children[r-'a']
		if node == nil {
			return false
		}
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, r := range prefix {
		node = node.children[r-'a']
		if node == nil {
			return false
		}
	}
	return true
}
