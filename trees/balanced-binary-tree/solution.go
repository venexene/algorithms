package main

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
 
func isBalanced(root *TreeNode) bool {
	res, _ := dfs(root)
	return res
}

func dfs(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}

	leftFlag, leftDepth := dfs(node.Left)
	rightFlag, rightDepth := dfs(node.Right)
	flag := max(leftDepth-rightDepth, rightDepth-leftDepth) <= 1

	return leftFlag && rightFlag && flag, max(leftDepth, rightDepth) + 1
}