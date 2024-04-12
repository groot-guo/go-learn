package main

import (
	"math/rand"
	"time"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 应该当改变思路，取中间值做根节点，减少高度
func sortedArrayToBST2(nums []int) *TreeNode {
	mid := len(nums) / 2
	node := &TreeNode{Val: nums[mid]}
	helperTreeNodeRight(nums[mid+1:], node)
	helperTreeNodeLeft(nums[:mid], node)
	return node
}

func helperTreeNodeRight(nums []int, node *TreeNode) {
	for i := 0; i < len(nums); i++ {
		data := &TreeNode{Val: nums[i]}
		node.Right = data
		node = node.Right
	}
}

func helperTreeNodeLeft(nums []int, node *TreeNode) {
	length := len(nums) - 1
	for ; length >= 0; length-- {
		data := &TreeNode{Val: nums[length]}
		node.Left = data
		node = node.Left
	}
}

// 中序遍历选择左节点
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return helperTreeNode(nums, 0, len(nums)-1)
}

func helperTreeNode(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = helperTreeNode(nums, left, mid-1)
	node.Right = helperTreeNode(nums, mid+1, right)
	return node
}

// 中序遍历选择右节点
func sortedArrayToBST33(nums []int) *TreeNode {
	return helper44(nums, 0, len(nums)-1)
}

func helper44(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	// 总是选择中间位置右边的数字作为根节点
	mid := (left + right + 1) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper44(nums, left, mid-1)
	root.Right = helper44(nums, mid+1, right)
	return root
}

// 中序遍历，随机选择一个节点
func sortedArrayToBST234(nums []int) *TreeNode {
	rand.Seed(time.Now().UnixNano())
	return helper234(nums, 0, len(nums)-1)
}

func helper234(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	// 选择任意一个中间位置数字作为根节点
	mid := (left + right + rand.Intn(2)) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper234(nums, left, mid-1)
	root.Right = helper234(nums, mid+1, right)
	return root
}
