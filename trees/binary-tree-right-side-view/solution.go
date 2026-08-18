package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	max := 0
	rightSideViewAdd(root, &res, 0, &max)
	return res
}

func rightSideViewAdd(node *TreeNode, view *[]int, depth int, max *int) {
	depth++
	if depth > *max {
		*view = append(*view, node.Val)
		*max = depth
	}

	if node.Right != nil {
		rightSideViewAdd(node.Right, view, depth, max)
	}

	if node.Left != nil {
		rightSideViewAdd(node.Left, view, depth, max)
	}
}
