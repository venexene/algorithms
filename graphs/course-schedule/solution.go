package main

func main() {

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	for _, p := range prerequisites {
		adj[p[1]] = append(adj[p[1]], p[0])
	}

	res := true
	visited := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			res = res && canFinishAdd(adj, visited, i)
		}
	}
	return res
}

func canFinishAdd(adj [][]int, visited []int, course int) bool {
	visited[course] = 1

	for _, c := range adj[course] {
		if visited[c] == 1 {
			return false
		}
		if visited[c] == 0 {
			if !canFinishAdd(adj, visited, c) {
				return false
			}
		}
	}

	visited[course] = 2
	return true
}
