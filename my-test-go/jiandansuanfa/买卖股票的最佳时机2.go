package main

func maxProfit(prices []int) int {
	res := 0
	min := prices[0]
	for index := range prices {
		if prices[index]-min > 0 {
			res += prices[index] - min
		}
		min = prices[index]
	}

	return res
}

func maxProfit2(prices []int) (ans int) {
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
