package main

import "fmt"

func isPalindrome(s string) bool {
	formatedString := []rune{}
	for _, char := range s {
		// 大写转小写
		if char >= 'A' && char <= 'Z' {
			char += 'a' - 'A'
			formatedString = append(formatedString, char)
		} else if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			// 小写直接合入
			formatedString = append(formatedString, char)
		}
	}
	s = string(formatedString)
	reS := reverseString(s)

	return s == reS
}

// 双指针反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
}
