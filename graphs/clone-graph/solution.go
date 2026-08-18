package main

import "fmt"

func main() {
	node := &Node{1, []*Node{{2, []*Node{}}}}
	result := cloneGraph(node)
	fmt.Println(result.Neighbors[0].Val)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := map[*Node]*Node{}
	return dfs(node, visited)
}

func dfs(node *Node, visited map[*Node]*Node) *Node {
	clone := &Node{node.Val, []*Node{}}
	visited[node] = clone
	for _, n := range node.Neighbors {
		if _, ok := visited[n]; ok != true {
			clone.Neighbors = append(clone.Neighbors, dfs(n, visited))
		} else {
			clone.Neighbors = append(clone.Neighbors, visited[n])
		}
	}
	return clone
}
