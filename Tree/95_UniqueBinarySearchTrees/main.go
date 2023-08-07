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

func generateTrees(n int) []*TreeNode {
	return buildBST(1, n)
}

func buildBST(start int, end int) (results []*TreeNode) {
	if start > end {
		return nil
	}
	if start == end {
		results = append(results, &TreeNode{start, nil, nil})
		return results
	}
	for i := start; i < end+1; i++ {
		left := buildBST(start, i-1)
		right := buildBST(i+1, end)
		if len(left) != 0 && len(right) != 0 {
			for _, nodeLeft := range left {
				for _, nodeRight := range right {
					results = append(results, &TreeNode{i, nodeLeft, nodeRight})
				}
			}
		}
		if len(left) == 0 {
			for _, nodeRight := range right {
				results = append(results, &TreeNode{i, nil, nodeRight})
			}
		}
		if len(right) == 0 {
			for _, nodeLeft := range left {
				results = append(results, &TreeNode{i, nodeLeft, nil})
			}
		}
	}
	return results
}

func main() {
	fmt.Println(generateTrees(3))
}
