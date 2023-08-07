package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	stringArray := strings.Split(s, "")
	stringMap := make(map[string]int)
	temp := 0
	max := 0
	index := 0
	for i, ss := range stringArray {
		if _, ok := stringMap[ss]; !ok {
			stringMap[ss] = i
		} else {
			temp = i - index
			flag := stringMap[ss]
			if flag >= index {
				index = flag + 1
			}
			stringMap[ss] = i
		}
		if temp > max {
			max = temp
		}
	}
	if val := len(stringArray) - index; val > max {
		max = val
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
