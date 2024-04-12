package main

// 1. 需要判断是否唯一，然后给数组对应位的值赋值，所以其实已经去重，使用 map 就会开辟新的空间
func removeDuplicates(nums []int) int {
	keyMap := make(map[int]bool)
	i := 0
	for index := range nums {
		if _, ok := keyMap[nums[index]]; !ok {
			keyMap[nums[index]] = true
			nums[i] = nums[index]
			i++
		}
	}
	return i
}

// 1.直接进行当前位 与 目前需要比较的下标位
func removeDuplicates2(nums []int) int {
	var length = 1 // 最后长度
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[length] = nums[i]
			length++
		}
	}
	return length
}
