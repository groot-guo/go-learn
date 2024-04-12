package medium

import (
	"leetcode/my-test/mid-suanfa"
	"math"
	"sort"
)

/*
	三数之和
*/

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

/*
	矩阵置零
*/

// 标记数组  空间 m+n， 时间 mn
func setZeroes(matrix [][]int) {
	rows := make([]bool, len(matrix))
	cells := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cells[j] = true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if rows[i] || cells[j] {
				matrix[i][j] = 0
			}

		}
	}
}

// 两个标记变量
func setZeroes2(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	row0, col0 := false, false

	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}

	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
			break
		}
	}

	// 更新 1 列以前
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 更新 1 列以后
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if row0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}

	if col0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}
}

// 一个标记变量
func setZeroes3(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	col0 := false
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
		}
		for j := 1; j < m; j++ {
			if r[j] == 0 {
				r[0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}

/*
	字母异位词
*/
// 排序
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}

	for _, v := range strs {
		s := []byte(v)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], v)
	}

	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}

	return ans
}

func groupAnagrams2(strs []string) [][]string {
	mp := make(map[[26]int][]string)

	for _, v := range strs {
		cnt := [26]int{}
		for _, s := range v {
			n := s - 'a'
			cnt[n]++
		}
		mp[cnt] = append(mp[cnt], v)
	}

	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}

	return ans
}

/*
无重复字符的最长子串
*/

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1， 相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，删除第一个字符
			// 缩小窗口
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = mid_suanfa.max(ans, rk-i+1)
	}

	return ans
}

/*
	最长回文子串
*/

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

/*
	递增的三元子序列
*/

// 双向遍历
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	leftMin := make([]int, n)
	leftMin[0] = nums[0]
	for i := 1; i < n; i++ {
		leftMin[i] = mid_suanfa.min(leftMin[i-1], nums[i])
	}
	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = mid_suanfa.max(rightMax[i+1], nums[i])
	}

	for i := 1; i < n-1; i++ {
		if nums[i] > leftMin[i-1] && nums[i] < rightMax[i+1] {
			return true
		}
	}

	return false
}

func increasingTriplet2(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	first, second := nums[0], math.MaxInt32
	for i := 1; i < n; i++ {
		num := nums[i]
		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}
	}

	return false
}
