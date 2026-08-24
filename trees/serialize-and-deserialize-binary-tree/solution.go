package main

import (
	"strconv"
	"strings"
)

func main() {
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	
	builder := strings.Builder{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	builder.WriteString(strconv.Itoa(root.Val))

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			builder.WriteByte('|')
			builder.WriteString(strconv.Itoa(node.Left.Val))
			queue = append(queue, node.Left)
		} else {
			builder.WriteByte('|')
			builder.WriteString("#")
		}

		if node.Right != nil {
			builder.WriteByte('|')
			builder.WriteString(strconv.Itoa(node.Right.Val))
			queue = append(queue, node.Right)
		} else {
			builder.WriteByte('|')
			builder.WriteString("#")
		}
	}

	return builder.String()
}

func (this *Codec) deserialize(data string) *TreeNode {    
	if data == "" {
		return nil
	}
	
	vals := strings.Split(data, "|")
	val, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val:val}

	queue := []*TreeNode{}
	queue = append(queue, root)
	i := 1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if vals[i] != "#" {
			val, _ := strconv.Atoi(vals[i])
			node.Left = &TreeNode{Val: val}
			queue = append(queue, node.Left)
		}
		i++

		if vals[i] != "#" {
			val, _ := strconv.Atoi(vals[i])
			node.Right = &TreeNode{Val: val}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}
