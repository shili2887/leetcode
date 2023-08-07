package main

import (
	"fmt"
)

// no extra space = O(n)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var result int
	y := x
	for y != 0 {
		result = result*10 + y%10
		y /= 10
	}
	return x == result
}

func main() {
	fmt.Println(isPalindrome(121))
}
