package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	sums := make([]int, len(nums)+1)
	sums[0] = 0
	for i := 0; i < len(nums); i++ {
		sums[i+1] = sums[i] + nums[i]
	}

	minLen := math.MaxInt
	for i := 1; i <= len(nums); i++ {
		targetTmp := target + sums[i-1]
		index := binaryFind(sums, targetTmp)
		if index < 0 {
			index = len(nums) + 1
		}
		if index <= len(nums) {
			if minLen > index-(i-1) {
				minLen = index - (i - 1)
			}
		}
	}

	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}

func binaryFind(nums []int, target int) int {
	mid, l, r := -1, 0, len(nums)-1
	for l < r {
		mid = (l + r) >> 1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if nums[l] >= target {
		return l
	} else {
		return -1
	}
}

func main() {
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(80, []int{10, 5, 13, 4, 8, 4, 5, 11, 14, 9, 16, 10, 20, 8}))
}
