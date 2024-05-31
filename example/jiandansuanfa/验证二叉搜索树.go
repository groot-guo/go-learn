package main

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Val <= root.Left.Val {
		return false
	}
	if root.Val >= root.Right.Val {
		return false
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}

func isValidBST3(root *TreeNode) bool {
	return helper2(root, math.MinInt64, math.MaxInt64)
}

func helper2(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper2(root.Left, lower, root.Val) && helper2(root.Right, root.Val, upper)
}

func isValidBST5(root *TreeNode) bool {
	var pre *TreeNode
	var travel func(node *TreeNode) bool
	travel = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		r := travel(node.Left)
		if pre != nil && node.Val <= pre.Val {
			return false
		}
		pre = node
		l := travel(node.Right)
		return l && r
	}
	return travel(root)
}
