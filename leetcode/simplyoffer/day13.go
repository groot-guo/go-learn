package simplyoffer

import "strings"

/*
21. 调整数组顺序使奇数位于偶数前面
[1,2,3,4]
[1,3,2,4]
*/
func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	left, right := 0, len(nums)-1
	for left < right {
		if nums[right]%2 == 0 {
			right--
			continue
		}
		if nums[left]%2 != 0 {
			left++
			continue
		}

		nums[left], nums[right] = nums[right], nums[left]
		right--
		left++
	}

	return nums
}

// 运行时间效率最高
func exchange2(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[right]%2 == 0 {
			right--
		}
		for left < right && nums[left]%2 != 0 {
			left++
		}

		nums[left], nums[right] = nums[right], nums[left]
		right--
		left++
	}

	return nums
}

/*
57. 和为 s 的两个数字
nums = [2,7,11,15], target = 9
*/
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	left, right := 0, len(nums)-1
	for left < right {
		num := target - nums[left]
		if num > nums[right] {
			left++
			//right = len(nums) - 1 // 可以去除， 说明最右端数，远远大于 target
		} else if num == nums[right] {
			return []int{nums[left], nums[right]}
		} else {
			right--
		}
	}

	return nil
}

// 耗时最小
func twoSum2(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	left, right := 0, len(nums)-1
	for nums[left]+nums[right] != target {
		for nums[left]+nums[right] > target {
			right--
		}
		for nums[left]+nums[right] < target {
			left++
		}
	}

	return []int{nums[left], nums[right]}
}

/*
	58.翻转单词顺序
	"hello world!"
	"world! hello"
*/

func reverseWords(s string) string {
	left, right := len(s)-1, len(s)-1

	data := []string{}
	for left >= 0 {
		for left >= 0 && s[left] != ' ' {
			left--
		}
		newStr := strings.ReplaceAll(s[left+1:right+1], " ", "")
		if newStr != "" {
			data = append(data, newStr)
		}

		right = left
		left = right - 1
	}

	return strings.Join(data, " ")
}
