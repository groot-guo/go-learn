package main

import "math"

func maxProfit11(prices []int) int {
	price, left, right := 0, 0, 1

	for right < len(prices) {
		if prices[left] > prices[right] {
			left = right
		}
		res := prices[right] - prices[left]
		if res > price {
			price = res
		}

		right++
	}
	return price
}

// 运行效率最快
func maxProfit22(prices []int) int {
	buy := math.MinInt32
	sell := 0
	for _, p := range prices {
		buy = max33(buy, -p)
		sell = max33(sell, buy+p)
	}
	return sell
}

// max 求最大值
func max33(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 消耗内存最小
func maxProfit44(prices []int) int {
	dp_i_0 := 0
	dp_i_1 := math.MinInt32

	for i := 0; i < len(prices); i++ {
		if dp_i_0 < dp_i_1+prices[i] {
			dp_i_0 = dp_i_1 + prices[i]
		}
		if dp_i_1 < -prices[i] {
			dp_i_1 = -prices[i]
		}
	}
	return dp_i_0
}
