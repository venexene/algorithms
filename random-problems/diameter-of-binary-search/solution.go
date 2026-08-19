package main

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
    dfs(root, &diameter)
	return diameter
}

func dfs(node *TreeNode, diameter *int) int {
	if node == nil {
		return 0
	}
	left := dfs(node.Left, diameter)
	right := dfs(node.Right, diameter)
	*diameter = max(left+right, *diameter)
	return max(left, right) + 1
}