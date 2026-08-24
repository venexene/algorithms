package main

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
		return nil
	}
	
	inorderMap := map[int]int{}
	for i, n := range inorder {
		inorderMap[n] = i	
	}

	idx := 0
	var dfs func(left, right int) *TreeNode
	dfs = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		rootVal := preorder[idx]
		idx++

		root := &TreeNode{Val:rootVal}
		mid := inorderMap[rootVal]

		root.Left = dfs(left, mid-1)
		root.Right = dfs(mid+1, right)
		return root
	}

	return dfs(0, len(inorder) -1)
}

func dfsOld(preorder []int, inorder map[int]int, i, left, right int, node *TreeNode) int {
	if i >= len(preorder) {
		return i
	}
	mid := inorder[node.Val]

	pos := inorder[preorder[i]]
	if left <= pos && pos < mid {
		node.Left = &TreeNode{Val: preorder[i]}
		i = dfsOld(preorder, inorder, i+1, left, mid-1, node.Left)
	}

	if i >= len(preorder) {
        return i
    }

	pos = inorder[preorder[i]]
	if mid < pos && pos <= right {
		node.Right = &TreeNode{Val: preorder[i]}
		i = dfsOld(preorder, inorder, i+1, mid+1, right, node.Right)
	}

	return i
}