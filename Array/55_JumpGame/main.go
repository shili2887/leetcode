package main

import (
	"fmt"
	"math"
)

func canJump(nums []int) bool {
	maxStep := 0
	for i := 0; i < len(nums); i++ {
		// 说明数组内存在不可达的位置，直接返回
		if i > maxStep {
			return false
		}

		maxStep = int(math.Max(float64(maxStep), float64(i+nums[i])))
		if maxStep > len(nums) {
			return true
		}
	}

	return true
}

func main() {

	nums := []int{8, 2, 4, 4, 4, 9, 5, 2, 5, 8, 8, 0, 8, 6, 9, 1, 1, 6, 3, 5, 1, 2, 6, 6, 0, 4, 8, 6, 0, 3, 2, 8, 7, 6, 5, 1, 7, 0, 3, 4, 8, 3, 5, 9, 0, 4, 0, 1, 0, 5, 9, 2, 0, 7, 0, 2, 1, 0, 8, 2, 5, 1, 2, 3, 9, 7, 4, 7, 0, 0, 1, 8, 5, 6, 7, 5, 1, 9, 9, 3, 5, 0, 7, 5}
	fmt.Println(canJump(nums))
}
