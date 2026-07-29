package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
		return false
	}

	return hasPathSumAdd(root, targetSum)
}

func hasPathSumAdd(root *TreeNode, targetSum int) bool {
    if root == nil {
		return targetSum == 0
	}

	targetSum -= root.Val 
	
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

