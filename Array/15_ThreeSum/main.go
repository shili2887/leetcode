package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	if len(nums) == 0 || nums[0] > 0 {
		return result
	}
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			start, end, sum := i+1, len(nums)-1, -nums[i]
			for start < end {
				if nums[start]+nums[end] == sum {
					result = append(result, []int{nums[i], nums[start], nums[end]})
					for start < end && nums[start] == nums[start+1] {
						start++
					}
					for start < end && nums[end] == nums[end-1] {
						end--
					}
					start++
					end--
				} else if nums[start]+nums[end] < sum {
					for start < end && nums[start] == nums[start+1] {
						start++
					}
					start++
				} else {
					for start < end && nums[end] == nums[end-1] {
						end--
					}
					end--
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{}))
}
