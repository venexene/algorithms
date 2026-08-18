package main

func main() {

}

func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	if n <= 0 {
		return false
	}

	visited := map[int]bool{} // можно использовать массив

	adj := map[int][]int{} // можно использовать массив
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	return validPathAdd(adj, source, destination, visited)
}

func validPathAdd(adj map[int][]int, source int, destination int, visited map[int]bool) bool {
	if source == destination {
		return true
	}

	visited[source] = true

	res := false
	for _, vertex := range adj[source] {
		if !visited[vertex] {
			res = res || validPathAdd(adj, vertex, destination, visited)
		}
	}
	return res
}

func validPathUnionFind(n int, edges [][]int, source int, destination int) bool {
	parent := map[int]int{} // можно использовать массив

	for i := 0; i < n; i++ {
		parent[i] = i
	}

	for _, edge := range edges {
		root1, root2 := find(edge[0], parent), find(edge[1], parent)
		parent[root2] = root1
	}

	return find(source, parent) == find(destination, parent)
}

func find(node int, parent map[int]int) int {
	if parent[node] != node {
		parent[node] = find(parent[node], parent)
	}
	return parent[node]
}
