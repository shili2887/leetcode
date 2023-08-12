package main

import "fmt"

// 双指针步长为2
func removeDuplicates(nums []int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if j < len(nums)-2 && nums[j] == nums[j+2] {
			continue
		}

		nums[i] = nums[j]
		i++
	}

	return i
}

func main() {
	nums := []int{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 3, 3, 3}
	fmt.Println(nums[:removeDuplicates(nums)])
}
