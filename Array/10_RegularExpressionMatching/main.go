package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	lenS := len(s)
	lenP := len(p)
	if lenP == 0 {
		return lenS == 0
	}
	if lenS == 0 {
		if lenP == 1 || p[1] != '*' {
			return false
		}
		return isMatch(s, p[2:])
	}
	switch p[0] {
	case '.':
		if lenP == 1 {
			return isMatch(s[1:], p[1:])
		}
		if p[1] == '*' {
			return isMatch(s[1:], p) || isMatch(s, p[2:])
		}
		return isMatch(s[1:], p[1:])
	case '*':
		return false
	default:

		if p[0] == s[0] {
			if lenP == 1 {
				return isMatch(s[1:], p[1:])
			}
			if p[1] == '*' {
				return isMatch(s[1:], p) || isMatch(s, p[2:])
			}
			return isMatch(s[1:], p[1:])
		}
		if lenP == 1 {
			return false
		}
		if p[1] == '*' {
			return isMatch(s, p[2:])
		}
		return false
	}
}

func main() {
	fmt.Println(isMatch("aab", "aa*c*b*a*b"))
}
