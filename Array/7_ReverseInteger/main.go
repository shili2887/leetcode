package main

import (
	"fmt"
	"math"
)

// 2147483647
func reverse(x int) int {
	var result int
	temp := int(math.Abs(float64(x)))
	for temp != 0 {
		mod := temp % 10
		result = result*10 + mod
		temp /= 10
	}
	if x < 0 {
		if result > 2147483648 {
			return 0
		}
		return 0 - result
	}
	if result > 2147483647 {
		return 0
	}
	return result

}

func main() {
	fmt.Printf("%d", reverse(120))
}
