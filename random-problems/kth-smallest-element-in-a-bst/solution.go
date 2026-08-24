package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return root
		}

		left := dfs(root.Left)
		if left != nil {
			return left
		}
		k -= 1
		if k == 0 {
			return root
		}
		return dfs(root.Right)
	}
	return dfs(root).Val
}
