package main

import "fmt"

func lengthOfLastWord(s string) int {
	end := len(s)
	for end > 0 && s[end-1] == ' ' {
		end--
	}

	if end < 0 {
		return 0
	}

	start := end
	for start > 0 && s[start-1] != ' ' {
		start--
	}

	return end - start
}

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
}
