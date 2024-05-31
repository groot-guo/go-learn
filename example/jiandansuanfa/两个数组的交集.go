package main

import "sort"

/*
	两个数组的交集 2
	[1,2,2,1]  [2,2]
	交集应当位 [2,2]

	如果给定的数组已经排好序呢？你将如何优化你的算法？
	如果nums1的大小比nums2 小，哪种方法更优？
	如果nums2的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

	1. 采用数组下标的方式。 数组的 1 的下标位置 与 数组 2 的下标开始比较, 如果需要考虑顺序
	2. 不需要考虑顺序，排序，之后再进行比较，
*/

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i := 0
	index := 0
	res := make([]int, 0, len(nums1))
	for i < len(nums1) && index < len(nums2) {
		if nums1[i] < nums2[index] {
			i++
		} else if nums1[i] > nums2[index] {
			index++
		} else if nums1[i] == nums2[index] {
			res = append(res, nums1[i])
			i++
			index++
		}
	}
	return res
}

// 消耗内存最小
func intersect2(nums1 []int, nums2 []int) []int {
	len2 := len(nums2)
a:
	for i := len(nums1) - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				if j != (len2 - 1) {
					nums2[len2-1] = nums2[j] ^ nums2[len2-1]
					nums2[j] = nums2[j] ^ nums2[len2-1]
					nums2[len2-1] = nums2[j] ^ nums2[len2-1]
				}
				len2--
				continue a
			}
		}
	}

	return nums2[len2:]
}
