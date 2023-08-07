package main

import (
	"fmt"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) rigthAppend(val int) {
	t.Right = &TreeNode{val, nil, nil}
}

func (t *TreeNode) leftAppend(val int) {
	t.Left = &TreeNode{val, nil, nil}
}

func postorderTraversal(root *TreeNode) []int {
	var cur int
	var stack []*TreeNode
	var result []int

	if root == nil {
		return result
	}
	stack = append(stack, root)
	for {
		for stack[cur].Left != nil {
			stack = append(stack, stack[cur].Left)
			stack[cur].Left = nil
			cur++
		}
		for stack[cur].Right != nil {
			stack = append(stack, stack[cur].Right)
			stack[cur].Right = nil
			cur++
		}
		result = append(result, stack[cur].Val)
		stack = stack[:cur]
		cur--
		if cur == -1 {
			return result
		}
	}
}

func inorderTraversal(root *TreeNode) []int {

	var cur int
	var stack []*TreeNode
	var result []int
	if root == nil {
		return result
	}
	stack = append(stack, root)
	for {
		for stack[cur].Left != nil {
			stack = append(stack, stack[cur].Left)
			stack[cur].Left = nil
			cur++
		}
		result = append(result, stack[cur].Val)
		if stack[cur].Right != nil {
			stack[cur] = stack[cur].Right
		} else {
			stack = stack[:cur]
			cur--
			if cur == -1 {
				return result
			}
		}
	}
}

func morrisTraversal(root *TreeNode) []int {
	var prev, cur *TreeNode
	var result []int
	cur = root
	for cur != nil {
		if cur.Left == nil {
			result = append(result, cur.Val)
			cur = cur.Right
		} else {
			prev = cur.Left
			for prev.Right != nil && prev.Right != cur {
				prev = prev.Right
			}
			if prev.Right == nil {
				prev.Right = cur
				cur = cur.Left
			} else {
				result = append(result, cur.Val)
				cur = cur.Right
			}
		}
	}
	return result
}

func main() {
	root := &TreeNode{1, nil, nil}
	root.rigthAppend(5)
	root.leftAppend(4)
	root.Left.leftAppend(4)
	root.Left.rigthAppend(4)
	root.Right.rigthAppend(5)
	fmt.Println(morrisTraversal(root))
}
