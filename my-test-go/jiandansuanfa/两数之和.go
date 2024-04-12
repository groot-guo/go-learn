package main

import "fmt"

/*
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。


*/

func main() {
	//data := []int{2, 7, 11, 15}
	data := []int{3, 2, 4}
	res := twoSum(data, 6)
	fmt.Println(res)
}

// 耗时较长，内存使用较低
func twoSum(nums []int, target int) []int {

	left := 0
	right := 1

	res := make([]int, 0)
	for left < len(nums) && right < len(nums) {
		if left == right {
			break
		}
		if nums[left]+nums[right] == target {
			res = append(res, left)
			res = append(res, right)
		}
		right++

		if right == len(nums) {
			left++
			right = left + 1
		}
	}

	return res
}

// 用时最快
func twoSum2(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for i, v := range nums {
		if p, ok := hashTable[target-v]; ok {
			return []int{p, i}
		}
		hashTable[v] = i
	}
	return nil
}
