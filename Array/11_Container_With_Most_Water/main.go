package main

import (
	"fmt"
)

func maxArea(height []int) int {
	var max, temp int
	length := len(height)
	start := 0
	end := length - 1
	for end > start {
		if height[start] > height[end] {
			temp = (end - start) * height[end]
			end--
		} else {
			temp = (end - start) * height[start]
			start++
		}
		if temp > max {
			max = temp
		}
	}
	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 1}))
}
