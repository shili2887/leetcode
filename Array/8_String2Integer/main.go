package main

import (
	"fmt"
)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}

	var index, result int
	term := len(str)
	array := []rune(str)
	flag := 1
	for ind, val := range array {
		if val != 32 {
			index = ind
			break
		}
	}

	if array[index] == '-' {
		flag = -1
		index++
	} else if array[index] == '+' {
		flag = 1
		index++
	}
	for ind, val := range array[index:] {
		if val < '0' || val > '9' {
			term = ind + index
			break
		}
	}
	if len(array[index:term]) > 10 {
		if flag > 0 {
			return 2147483647
		}
		return -2147483648
	}

	for i := index; i < term; i++ {
		if array[i] > 47 && array[i] < 58 {
			result = result*10 + int(array[i]) - 48
		}
	}
	if flag*result < 0 {
		result = -result
	}
	if result > 2147483647 {
		return 2147483647
	}
	if result < -2147483648 {
		return -2147483648
	}
	return result
}

func main() {
	fmt.Println(myAtoi("+-2"))
}
