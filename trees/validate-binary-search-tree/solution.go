package main

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
    if root == nil {
		return true
	}

	min := math.MinInt64
	max := math.MaxInt64
	return isValidBSTAdd(root, min, max)
}

func isValidBSTAdd(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	left := isValidBSTAdd(node.Left, min, node.Val)

	right := isValidBSTAdd(node.Right, node.Val, max)

	return left && right
}