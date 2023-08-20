package main

import (
	"fmt"
	"sort"
)

// 备忘录，什么破题干，写完了才看懂。。
func hIndex(citations []int) int {
	nums := make([]int, len(citations)+1)
	hindex := 0

	for _, citation := range citations {
		if citation > len(citations) {
			nums[len(citations)]++
		} else {
			nums[citation]++
		}
	}

	for i := len(citations); i >= 0; i-- {
		hindex += nums[i]
		if hindex >= i {
			return i
		}
	}

	return hindex
}

// 排序法
func hIndexSort(citations []int) int {
	sort.SliceStable(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	hindex := 0
	for index, citation := range citations {
		if index+1 <= citation {
			hindex = index + 1
		} else {
			break
		}
	}

	return hindex
}

func main() {
	nums := []int{0}
	fmt.Println(hIndex(nums))
}
