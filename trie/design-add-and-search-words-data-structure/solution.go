package main

func main() {

}

type WordDictionary struct {
	root *Node
}

type Node struct {
	children [26]*Node
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &Node{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, r := range word {
		if node.children[r-'a'] == nil {
			node.children[r-'a'] = &Node{}
		}
		node = node.children[r-'a']
	}
	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.SearchDFS(word, this.root)
}

func (this *WordDictionary) SearchDFS(word string, node *Node) bool {
	for i, r := range word {
		if r == '.' {
			res := false
			for _, child := range node.children {
				if child != nil {
					res = res || this.SearchDFS(word[i+1:], child)
				}
			}
			return res
		}
		node = node.children[r-'a']
		if node == nil {
			return false
		}
	}
	return node.isEnd
}
