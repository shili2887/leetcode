package main

import (
	"fmt"
	"math"
)

// 递归备忘录，如果有结果，直接返回
var RecurResult = make(map[int]int)

// 递归+备忘录
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	// 由于题干保证一定可以到最后，那么长度小于3的话一定是1
	if len(nums) < 3 {
		return 1
	}
	return jumpR(nums, 0)
}

func jumpR(nums []int, index int) int {
	if index >= len(nums)-1 {
		return 0
	}

	// 如果当前值是0，说明永远到不了最后，返回-1
	if nums[index] == 0 {
		return -1
	}

	length := nums[index]
	minStep := math.MaxInt
	for i := 1; i <= length; i++ {
		tmp := 0
		if tmpMem, ok := RecurResult[index+i]; !ok {
			tmp = jumpR(nums, index+i)
			RecurResult[index+i] = tmp
		} else {
			tmp = tmpMem
		}
		if tmp == -1 {
			continue
		}

		if tmp < minStep {
			minStep = tmp
		}
	}

	// 如果最后这跳链路不能跳转，返回-1
	if minStep == math.MaxInt {
		return -1
	}

	return minStep + 1
}

// DP
func jumpDP(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		return 1
	}

	maxStep := 1
	dp := make([]int, len(nums))
	for i := 1; i < len(nums)+1; i++ {
		num := nums[i-1]

		minStep := dp[i-1] + 1
		if i+num >= len(nums) {
			return minStep
		}
		for j := maxStep; j < i+num; j++ {
			if dp[j] == 0 {
				dp[j] = minStep
			}
		}
		if maxStep < i+num {
			maxStep = i + num
		}
	}

	return -1
}

func main() {
	nums := []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}
	fmt.Println(jumpDP(nums))
}
