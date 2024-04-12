package main

import "sort"

/*
	输入一个字符串，然后输出最长有效括号的长度
	仅包含左右括号

	()()
	()(()
	((()))
*/

func get(s string) int {
	ans := 0

	stack := make([]int, 0, len(s))
	stack = append(stack, -1)
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				ans = max(ans, i-stack[len(stack)-1])
			}
		}
	}
	i := sort.SearchInts(len(nums[right:]), func(n int) bool {
		return nums[right:][n] == num1
	})

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
