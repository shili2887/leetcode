package main

import (
	"fmt"
	"math"
)

func longestPalindrome(s string) string {
	array := []rune(s)
	maxLen := 0
	index := 0
	for i := 0; i < len(array); i++ {
		maxTemp := -1
		var indexTemp int
		left := i
		right := i
		for j := 0; j < int(math.Min(float64(len(array)-i-1), float64(i))+1); j++ {
			if array[left-j] == array[right+j] {
				maxTemp += 2
				indexTemp = left - j
			} else {
				break
			}
		}
		if maxTemp > maxLen {
			index = indexTemp
			maxLen = maxTemp
		}

		maxTemp = 0
		right++
		if right < len(array) {
			for j := 0; j < int(math.Min(float64(len(array)-i-2), float64(i))+1); j++ {
				if array[left-j] == array[right+j] {
					maxTemp += 2
					indexTemp = left - j
				} else {
					break
				}
			}
			if maxTemp > maxLen {
				index = indexTemp
				maxLen = maxTemp
			}
		}
	}
	return string(array[index : index+maxLen])
}

func main() {
	fmt.Println(longestPalindrome("cccaccc"))
}
