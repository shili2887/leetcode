package main

import (
	"fmt"
)

type anR struct {
	Num int
	Rom string
}

func getRom() (array [13]anR) {

	array[0] = anR{1000, "M"}
	array[1] = anR{900, "CM"}
	array[2] = anR{500, "D"}
	array[3] = anR{400, "CD"}
	array[4] = anR{100, "C"}
	array[5] = anR{90, "XC"}
	array[6] = anR{50, "L"}
	array[7] = anR{40, "XL"}
	array[8] = anR{10, "X"}
	array[9] = anR{9, "IX"}
	array[10] = anR{5, "V"}
	array[11] = anR{4, "IV"}
	array[12] = anR{1, "I"}

	return array
}

func romanToInt(s string) int {
	var result int
	array := getRom()
	for index := 0; len(s) > 0; {
		rome := array[index].Rom
		length := len(rome)
		if len(s) < length || string(s[:length]) != rome {
			index++
		} else {
			s = s[length:]
			result += array[index].Num
		}
	}
	return result
}

func main() {
	fmt.Println(romanToInt("DCXXI"))
}
