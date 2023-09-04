package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 1 {
		return false
	}
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == len(s)
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
