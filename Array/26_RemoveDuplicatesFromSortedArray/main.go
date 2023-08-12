package main

import "fmt"

// 快排逻辑
func removeDuplicates(nums []int) int {
	i := 1
	flag := nums[0]
	for i < len(nums) {
		if nums[i] > flag {
			flag = nums[i]
			i++
		} else {
			j := i + 1
			for j < len(nums) {
				if nums[j] <= flag {
					j++
					continue
				}
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
				break
			}

			if j == len(nums) {
				break
			}
		}
	}

	return i
}

// 双指针简单逻辑
func removeDuplicatesSimp(nums []int) int {
	j := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}

func main() {
	nums := []int{1, 2}

	fmt.Println(removeDuplicatesSimp(nums))
}
