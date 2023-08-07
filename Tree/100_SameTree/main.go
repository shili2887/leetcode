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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil {
		return false
	}
	if q == nil && p != nil {
		return false
	}
	if q == nil && p == nil {
		return true
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	p := &TreeNode{1, nil, nil}
	q := &TreeNode{1, nil, nil}
	p.leftAppend(2)
	q.rigthAppend(2)
	fmt.Println(isSameTree(p, q))
}
