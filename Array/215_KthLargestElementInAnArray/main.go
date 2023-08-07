package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	right := len(nums) - 1
	left := 0
	flag := nums[left]
	i, j := left, right
	p := 0
	for j >= i {
		for j >= p && nums[j] <= flag {
			j--
		}
		if j >= p {
			nums[p] = nums[j]
			p = j
		}
		for i <= p && nums[i] >= flag {
			i++
		}
		if i <= p {
			nums[p] = nums[i]
			p = i
		}
	}
	nums[p] = flag
	switch {
	case p == (k - 1):
		return nums[p]
	case p > (k - 1):
		return findKthLargest(nums[:p], k)
	case p < (k - 1):
		return findKthLargest(nums[p+1:], (k - p - 1))
	}
	return nums[k]
}

func main() {
	ans := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	fmt.Println(ans)
}
