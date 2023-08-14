package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minPrice := -1
	totalPrice := 0
	for index, price := range prices {
		// 波谷时买入
		if minPrice == -1 && index < len(prices)-1 && ((index == 0 && price <= prices[index+1]) || (price < prices[index+1] && price <= prices[index-1])) {
			minPrice = price
		}

		// 波峰时卖出
		if index > 0 && minPrice != -1 && ((index == len(prices)-1 && price >= prices[index-1]) || (price >= prices[index-1] && price > prices[index+1])) {
			totalPrice += price - minPrice
			minPrice = -1
		}
	}

	return totalPrice
}

// 动态规划
func maxProfitDp(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	dp := [][2]int{}
	dp = append(dp, [2]int{0, -prices[0]})
	for i := 1; i < len(prices); i++ {
		dp = append(dp, [2]int{int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]+prices[i]))), int(math.Max(float64(dp[i-1][1]), float64(dp[i-1][0]-prices[i])))})
	}

	return dp[len(prices)-1][0]
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	nums1 := []int{7, 1, 5, 3, 6, 4}
	nums2 := []int{7, 6, 4, 3, 1}
	nums3 := []int{2, 2, 5}
	nums4 := []int{0, 5, 5, 6, 2, 1, 1, 3}
	nums5 := []int{5, 2, 3, 2, 6, 6, 2, 9, 1, 0, 7, 4, 5, 0}
	fmt.Println(maxProfit(nums))
	fmt.Println(maxProfit(nums1))
	fmt.Println(maxProfit(nums2))
	fmt.Println(maxProfit(nums3))
	fmt.Println(maxProfit(nums4))
	fmt.Println(maxProfitDp(nums5))
}
