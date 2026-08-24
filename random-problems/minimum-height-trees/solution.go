package main

func main() {

}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	tree := make([][]int, n)
	degree := make([]int, n)
	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
		degree[edge[0]]++
		degree[edge[1]]++
	}

	queue := []int{}
	for node := range tree {
		if degree[node] == 1 {
			queue = append(queue, node)
		}
	}

	remaining := n
	for remaining > 2 {
		layerSize := len(queue)
		remaining -= layerSize

		nextQueue := []int{}
		for i := 0; i < layerSize; i++ {
			leaf := queue[i]
			for _, neighbour := range tree[leaf] {
				degree[neighbour]--
				if degree[neighbour] == 1 {
					nextQueue = append(nextQueue, neighbour)
				}
			}
		}
		queue = nextQueue
	}

	return queue
}
