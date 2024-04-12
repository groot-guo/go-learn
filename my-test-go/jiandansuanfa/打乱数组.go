package main

import "math/rand"

// 存两个数组

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	res := make([]int, len(this.nums))
	copy(res, this.nums)
	rand.Shuffle(len(this.nums), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
