package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxPrice := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxPrice {
			maxPrice = price - minPrice
		}
	}

	return maxPrice
}

func main() {
	nums := []int{2, 4, 1}
	fmt.Println(maxProfit(nums))
}
