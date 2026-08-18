package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	depth := 0
	res := depth

	if root == nil {
		return 0
	}

	res = maxDepthRec(root, depth)

	return res
}

func maxDepthRec(node *TreeNode, depth int) int {
	depth++
	res := depth

	if node.Left != nil {
		res = max(res, maxDepthRec(node.Left, depth))
	}

	if node.Right != nil {
		res = max(res, maxDepthRec(node.Right, depth))
	}

	return res
}

func maxDepthShort(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
