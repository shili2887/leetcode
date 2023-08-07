package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var p *ListNode
	var headNode ListNode
	head = &headNode
	p = head
	var sum int
	for carrier := 0; l1 != nil || l2 != nil || carrier == 1; {
		var v1, v2 int
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
		}
		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
		}
		sum = v1 + v2 + carrier

		if sum > 9 {
			p.Val = sum - 10
			carrier = 1
		} else {
			p.Val = sum
			carrier = 0
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil && carrier == 0 {
			return head
		}
		p.Next = &ListNode{0, nil}
		p = p.Next
	}
	return head
}
