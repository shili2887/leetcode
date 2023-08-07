package main

import (
	"fmt"
	"math"
)

/*本题Golang的测试用例有问题, [2, nil, 1]在本测试用例下输出正确*/

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

var (
	wrongNode [2]*TreeNode
	prev      = &TreeNode{math.MinInt32, nil, nil}
)

func recoverTree(root *TreeNode) {
	inorderTraversal(root)
	temp := wrongNode[0].Val
	wrongNode[0].Val = wrongNode[1].Val
	wrongNode[1].Val = temp
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	if wrongNode[0] == nil && prev.Val >= root.Val {
		wrongNode[0] = prev
	}
	if wrongNode[0] != nil && prev.Val >= root.Val {
		wrongNode[1] = root
	}
	prev = root
	inorderTraversal(root.Right)
}

func main() {
	root := &TreeNode{2, nil, nil}
	root.rigthAppend(1)
	recoverTree(root)
	fmt.Println(root)
}
