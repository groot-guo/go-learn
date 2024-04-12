package main

import "fmt"

func twoSum(nums []int, target int) (ans []int) {
	existedNumMap := make(map[int]int)

	for i := range nums {
		if _, ok := existedNumMap[nums[i]]; !ok {
			existedNumMap[nums[i]] = i
		}
	}

	for i := range nums {
		num1 := nums[i]
		num2 := target - num1
		index, ok := existedNumMap[num2]
		if ok {
			ans = make([]int, 2)
			ans[0], ans[1] = i, index
			return ans
		}
	}

	return nil
}

type MyStruct struct {
	value int
}

func main() {
	// 创建一个MyStruct实例并分配内存
	ms := &MyStruct{value: 10}

	// 将ms的地址赋值给一个新变量ptr
	ptr := ms

	// 将ms的内存释放，模拟垃圾回收
	ms = nil

	// 此时ptr指向的内存已经被释放，试图访问会导致不确定的行为（悬挂指针）
	fmt.Println(ptr.value) // 这是一个潜在的悬挂指针访问，运行时可能会崩溃或返回不确定的结果
}
