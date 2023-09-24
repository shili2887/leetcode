package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	noteMap := make(map[rune]int)
	for _, note := range magazine {
		noteMap[note]++
	}

	for _, note := range ransomNote {
		noteMap[note]--
		if noteMap[note] < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canConstruct("aa", "aab"))
}
