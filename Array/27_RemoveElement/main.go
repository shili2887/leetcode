package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] != val {
			continue
		}

		j := i + 1
		for ; j < len(nums); j++ {
			if nums[j] == val {
				continue
			}

			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			break
		}

		fmt.Println(nums)
		if j == len(nums) {
			break
		}
	}

	nums = nums[:i]
	return i
}

// 双指针优化
func removeElementTail(nums []int, val int) int {
	i := 0
	j := len(nums)
	for i < j {
		fmt.Println(i, j)
		if nums[i] == val {
			nums[i] = nums[j-1]
			j--
		} else {
			i++
		}
	}

	nums = nums[:j]
	fmt.Println(nums)
	return j
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}

	fmt.Println(removeElementTail(nums, 2))
}
