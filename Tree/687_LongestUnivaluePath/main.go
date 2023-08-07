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

func longestUnivaluePath(root *TreeNode) int {
	max := 0
	findPath(root, &max)
	return max
}

func findPath(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := findPath(root.Left, max)
	right := findPath(root.Right, max)

	var leftCount, rightCount int
	if root.Left != nil {
		if root.Val == root.Left.Val {
			leftCount = left + 1
		}
	}
	if root.Right != nil {
		if root.Val == root.Right.Val {
			rightCount = right + 1
		}
	}
	cur := leftCount + rightCount
	if *max < cur {
		*max = cur
	}

	res := leftCount
	if leftCount < rightCount {
		res = rightCount
	}
	return res
}

func main() {
	root := &TreeNode{1, nil, nil}
	root.rigthAppend(1)
	root.Right.leftAppend(1)
	root.Right.rigthAppend(1)
	root.Right.Left.leftAppend(1)
	root.Right.Left.rigthAppend(1)
	root.Right.Right.leftAppend(1)

	fmt.Println(longestUnivaluePath(root))
}
