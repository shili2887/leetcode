package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for j := 0; j < len(nums); j++ {
		m[nums[j]] = j
	}
	for i := 0; i < len(nums); i++ {
		val := target - nums[i]
		if res, ok := m[val]; ok && res != i {
			return []int{i, res}
		}
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}
