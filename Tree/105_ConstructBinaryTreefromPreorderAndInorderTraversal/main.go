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

func buildTree(preorder []int, inorder []int) *TreeNode {
	result := build(preorder, inorder)
	return result
}

func build(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var index int
	for key, val := range inorder {
		if val == preorder[0] {
			index = key
			break
		}
	}
	root := &TreeNode{preorder[0], nil, nil}
	root.Left = build(preorder[1:index+1], inorder[:index])
	root.Right = build(preorder[index+1:], inorder[index+1:])
	return root
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	result := buildTree(preorder, inorder)
	fmt.Println(result)
}
