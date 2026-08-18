package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBufQueue(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	bufQueue := []*TreeNode{}
	res := [][]int{}
	levelElems := []int{}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			bufQueue = append(bufQueue, node.Left)
		}

		if node.Right != nil {
			bufQueue = append(bufQueue, node.Right)
		}

		levelElems = append(levelElems, node.Val)
		if len(queue) == 0 {
			res = append(res, levelElems)
			levelElems = []int{}
			queue = bufQueue
			bufQueue = []*TreeNode{}
		}
	}

	return res
}

func levelOrderLen(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	res := [][]int{}

	for len(queue) != 0 {
		levelSize := len(queue)
		res = append(res, []int{})
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			res[len(res)-1] = append(res[len(res)-1], node.Val)
		}

	}

	return res
}
