package mid_suanfa

import (
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	sort.Ints(coins)

	res := 0
	n := len(coins) - 1
	l := n
	price := amount
	for price > 0 && l > 0 {
		n = l
		for n >= 0 {
			if price-coins[n] > 0 {
				price = price - coins[n]
				res += 1
				n--
			} else if price-coins[n] < 0 {
				n--
			} else {
				break
			}
		}
		l--
	}

	if res != 0 {
		return res
	}
	return -1
}

func coinChange555(coins []int, amount int) int {
	var dp [1e4 + 1]int
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, coin := range coins {
			if i-coin < 0 || dp[i-coin] == -1 {
				continue
			}
			if cur := dp[i-coin] + 1; dp[i] == -1 || dp[i] > cur {
				dp[i] = cur
			}
		}
	}
	return dp[amount]
}

func coinChange333(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i < amount+1; i++ {
		dp[i] = math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
