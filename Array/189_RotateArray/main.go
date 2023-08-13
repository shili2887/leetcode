package main

import (
	"fmt"
	"sort"
)

type linkNode struct {
	Num  int
	Next *linkNode
}

func (t *linkNode) append(val int) {
	t.Next = &linkNode{val, nil}
}

func buildLink(nums []int) *linkNode {
	var head, next *linkNode
	next = &linkNode{
		Num: nums[0],
	}
	head = next
	for i := 1; i < len(nums); i++ {
		next.Next = &linkNode{
			Num: nums[i],
		}
		next = next.Next
	}

	return head
}

func turnLinkedList2Array(head *linkNode) []int {
	nums := []int{}
	if head == nil {
		return nums
	}

	for head != nil {
		nums = append(nums, head.Num)
		head = head.Next
	}

	return nums
}

func rotate(nums []int, k int) {
	if len(nums) == 0 || len(nums) == 1 || len(nums) < k {
		return
	}
	head := buildLink(nums)
	next := head
	tail := head
	for i := 0; i < len(nums)-1; i++ {
		if i < len(nums)-k-1 {
			next = next.Next
		}
		tail = tail.Next
	}
	fmt.Println(next.Next.Num)

	tail.Next = head
	head = next.Next
	next.Next = nil

	nums = turnLinkedList2Array(head)
	fmt.Println(nums)
}

// 反转法
func rotateReverse(nums []int, k int) {
	k %= len(nums)
	if len(nums) == 0 || len(nums) == 1 || len(nums) < k {
		return
	}
	sort.SliceStable(nums, func(i, j int) bool { return i > j })
	sort.SliceStable(nums[:k], func(i, j int) bool { return i > j })
	sort.SliceStable(nums[k+1:], func(i, j int) bool { return i > j })

	fmt.Println(nums)
}

func main() {
	nums := []int{-1, -100, 3, 99}
	rotateReverse(nums, 2)

}
