package main

/*
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

如果数组中的所有元素都是9，类似9999，加1之后肯定会变为10000，也就是数组长度会增加1位。
如果数组的元素只要有一个不是9，加1之后直接返回即可。
*/

func plusOne(digits []int) []int {
	index := len(digits) - 1
	for index >= 0 {
		if digits[index] != 9 {
			digits[index] += 1
			return digits
		} else {
			digits[index] = 0
		}

		index--
	}

	res := make([]int, 0, len(digits)+1)
	res = append(res, 1)
	res = append(res, digits...)
	return res
}
