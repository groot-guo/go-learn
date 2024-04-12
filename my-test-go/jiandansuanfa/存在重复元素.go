package main

// 使用map 可以确定
func containsDuplicate(nums []int) bool {
	keyMap := make(map[int]bool)
	for index := range nums {
		if _, ok := keyMap[nums[index]]; ok {
			return true
		}
		keyMap[nums[index]] = true
	}
	return false
}

// 使用 排序，然后再进行前几位数判断比较
