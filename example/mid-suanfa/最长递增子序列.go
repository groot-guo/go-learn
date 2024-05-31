package mid_suanfa

import "sort"

func lengthOfLIS(nums []int) int {
	var dp = make([]int, len(nums))
	dp[0] = 1
	var ret = 1
	for i := 1; i < len(nums); i++ {
		var max = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && max < dp[j]+1 {
				max = dp[j] + 1
			}
		}
		dp[i] = max
		if ret < max {
			ret = max
		}
	}
	return ret
}

func lengthOfLIS66(nums []int) int {
	var g, p = make([]int, len(nums)), 0 //1)贪心
	for _, num := range nums {
		f := sort.SearchInts(g[:p], num)
		g[f] = num
		if f == p {
			p++
		}
	}
	return p
}
