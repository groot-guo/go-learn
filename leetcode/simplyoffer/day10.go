package simplyoffer

import "strconv"

/*
46.把数字翻译成字符串
动态规划解决思路
[1,4,0,2]
[14,0,2] 合格
[1, 40, 2] 不合法
[1, 4, 02] 不合法
字符串第 i 位置：1. 可以单独作为一位翻译， 2。 i-1 和 i 组成在 10 ~ 25 之间，可以翻译
单独翻译对 f(i) 的贡献是 f(i-1)
i-1 和 i 组合对 f(i) 的贡献就是 f(i-2)
f(i) = f(i-1) + f(i-2)
*/
func translateNum(num int) int {
	numStr := strconv.Itoa(num)

	p, q, r := 0, 0, 1
	for i := 0; i < len(numStr); i++ {
		p, q, r = q, r, 0
		r += q
		if i == 0 {
			continue
		}

		pre := numStr[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}

	return r
}

/*
48.最长不含重复字符的子字符串
pwwkew

i = 0, m = {p, w}, rk = 1, ans = 2
i = 1, m = {w}, rk = 1, ans = 2
i = 2, m = {w, k ,e}, rk = 4, ans = 3
*/
func lengthOfLongestSubstring(s string) int {
	m, n := map[byte]int{}, len(s)

	rk, ans := -1, 0
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(m, s[i-1])
		}

		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}

		ans = max(ans, rk-i+1)
	}

	return ans
}
