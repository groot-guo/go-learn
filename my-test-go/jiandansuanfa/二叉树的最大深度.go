package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftLen := maxDepth(root.Left) + 1
	rightLen := maxDepth(root.Right) + 1

	if leftLen < rightLen {
		return rightLen
	} else {
		return leftLen
	}
}

// 运行效率最快
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	cur, left, right := 1, 0, 0
	if root.Left != nil {
		left = maxDepth(root.Left)
	}
	if root.Right != nil {
		right = maxDepth(root.Right)
	}
	cur += max2(left, right)
	return cur
}

func max2(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// bfs 实现方式
// 层层遍历， 广度优先法则
func maxDepth44(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}

// dfs， 使用两个栈，分别记录节点的 stack 栈， 节点所在层数的 level 栈 深度优先
func maxDepth33(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max2(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
