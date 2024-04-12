package main

import "fmt"

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。

进阶：你能尽量减少完成的操作次数吗？


*/

func moveZeroes(nums []int) {
	res := make([]int, 0, len(nums))
	res = append(res, nums...)
	right := len(nums) - 1
	left := 0
	for i := 0; i < len(res) && left <= right; i++ {
		if res[i] == 0 {
			nums[right] = res[i]
			right--
		} else {
			nums[left] = res[i]
			left++
		}
	}
}

// 不复制原数组得情况下操作

func main() {
	data := []int{0, 1, 0, 3, 12}
	//data := []int{0, 0, 1}
	moveZeroes2(data)
	fmt.Println(data)
}

func moveZeroes2(nums []int) {
	right := 1
	left := 0
	for left < len(nums) && right < len(nums) {
		if nums[left] == 0 && nums[right] != 0 {
			nums[right], nums[left] = nums[left], nums[right]
			right++
			left++
		} else if nums[left] == 0 && nums[right] == 0 {
			right++
		} else {
			left++
			right++
		}
	}
}

// 执行快速
func moveZeroes3(nums []int) {
	news := make([]int, len(nums))
	tag := 0
	for _, v := range nums {
		if v != 0 {
			news[tag] = v
			tag++
		}
	}
	copy(nums, news)
}
