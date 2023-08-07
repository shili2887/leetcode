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

func zigzagLevelOrder(root *TreeNode) [][]int {
	var results [][]int
	traversal(root, 0, &results)
	return results
}

func traversal(root *TreeNode, level int, results *[][]int) {
	if root == nil {
		return
	}
	if len(*results) == level {
		temp := []int{root.Val}
		*results = append(*results, temp)
	} else {
		if level%2 == 0 {
			(*results)[level] = append((*results)[level], root.Val)
		} else {
			(*results)[level] = append([]int{root.Val}, (*results)[level]...)
		}
	}
	traversal(root.Left, level+1, results)
	traversal(root.Right, level+1, results)
}

func main() {
	root := &TreeNode{3, nil, nil}
	root.leftAppend(9)
	root.rigthAppend(20)
	root.Right.leftAppend(15)
	root.Right.rigthAppend(7)
	root.Left.leftAppend(5)
	root.Left.rigthAppend(17)
	fmt.Println(zigzagLevelOrder(root))
}
