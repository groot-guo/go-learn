package main

// 1. 使用临时数组， 因为超过数组长度从头开始，所以可以使用 (i+k)%len
func rotate(nums []int, k int) {
	res := make([]int, 0, len(nums))
	res = append(res, nums...)

	for i := 0; i < len(nums); i++ {
		nums[(i+k)%len(nums)] = res[i]
	}

}

// 分别反转前后
func rotate2(nums []int, k int) {
	lenth := len(nums)
	t := k % lenth
	reverse(nums)
	reverse(nums[:t])
	reverse(nums[t:])

}
func reverse(nums []int) {
	lenth := len(nums)
	for i := 0; i < lenth/2; i++ {
		nums[i], nums[lenth-i-1] = nums[lenth-i-1], nums[i]
	}
}
