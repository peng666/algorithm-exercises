package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	ans := 0
	minPrice := prices[0]
	for _, price := range prices {
		if price-minPrice > ans {
			ans = price - minPrice
		}
		if price < minPrice {
			minPrice = price
		}
	}
	return ans
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	ans := maxProfit(prices)
	fmt.Println(ans)
}
