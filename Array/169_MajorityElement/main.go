package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	numMap := make(map[int]int)

	maj := len(nums) / 2
	for _, num := range nums {
		numMap[num]++
	}

	for num, count := range numMap {
		if count > maj {
			return num
		}
	}

	return nums[0]
}

// 排序法
func majorityElementSort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 递归
func majorityElementD(nums []int) int {
	return majorityElementDD(nums, 0, len(nums)-1)
}

// 递归子函数
func majorityElementDD(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}

	mid := (start + end) / 2
	left := majorityElementDD(nums, start, mid)
	right := majorityElementDD(nums, mid+1, end)

	var leftNum, rightNum int
	for i := start; i <= end; i++ {
		if nums[i] == left {
			leftNum++
		}
		if nums[i] == right {
			rightNum++
		}
	}
	if leftNum > rightNum {
		return left
	} else {
		return right
	}
}

// Moore Vote
func majorityElementMoore(nums []int) int {
	var cand, count int
	for _, num := range nums {
		if count == 0 {
			cand = num
		}
		if cand == num {
			count++
		} else {
			count--
		}
	}

	return cand
}

func main() {
	nums := []int{6, 5, 5}
	fmt.Println(majorityElementMoore(nums))
}
