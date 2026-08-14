package main

import (
	"fmt"
	"strings"
)

func main() {
	dictionary := []string{"a","b","c"}
	sentence := "aadsfasf absbs bbab cadsfafs"
	result := replaceWords(dictionary, sentence)
	fmt.Println(result)
}

type Trie struct {
    root *Node
}

type Node struct {
	children [26]*Node
	isEnd bool
}

func Constructor() Trie {
	return Trie{
		root: &Node{},
	}
}

func (this *Trie) Insert(word string)  {
	node := this.root
	for _, r := range word {
		if node.children[r-'a'] == nil {
			node.children[r-'a'] = &Node{}
		}
		node = node.children[r-'a']
	}
	node.isEnd = true
}


func (this *Trie) FindRoot(word string) string {
	node := this.root
	var sb strings.Builder
	for _, r := range word {
		node = node.children[r-'a']
		if node == nil {
			return ""
		}
		sb.WriteByte(byte(r))
		if node.isEnd {
			return sb.String()
		}
	}
	return ""
}

func replaceWords(dictionary []string, sentence string) string {
    trie := Constructor()

	for _, word := range dictionary {
		trie.Insert(word)
	}
	
	words := strings.Split(sentence, " ")

	for i := 0; i < len(words); i++ {
		root := trie.FindRoot(words[i])
		if root != "" {
			words[i] = root
		}
	}

	return strings.Join(words, " ")
}