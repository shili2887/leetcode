package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	var numCols int
	array := []rune(s)
	length := len(array)
	newArray := make([][]rune, numRows)
	gap := numRows*2 - 2
	if gap == 0 {
		numCols = len(array)
	} else {
		numCols = length/gap + 1
	}

	index := 0
loop:
	for j := 0; j < numCols; j++ {
		for i := 0; i < numRows; i++ {
			if index == length {
				break loop
			}
			newArray[i] = append(newArray[i], array[index])
			index++
		}
		for i := numRows - 2; i > 0; i-- {
			if index == length {
				break loop
			}
			newArray[i] = append(newArray[i], array[index])
			index++
		}
	}
	var resultArray []rune
	for _, strings := range newArray {
		for _, str := range strings {
			resultArray = append(resultArray, str)
		}
	}
	return string(resultArray)
}

func main() {
	str := convert("AB", 1)
	fmt.Println(str)
}
