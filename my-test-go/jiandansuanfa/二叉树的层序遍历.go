package main

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var data [][]int

	helperLevel(&data, root, 0)

	return data
}

func helperLevel(data *[][]int, node *TreeNode, level int) {
	if node == nil {
		return
	}

	if level >= len(*data) {
		*data = append(*data, []int{})
	}

	(*data)[level] = append((*data)[level], node.Val)

	helperLevel(data, node.Left, level+1)
	helperLevel(data, node.Right, level+1)
}

// 用时最快
func levelOrder4(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := list.New()
	queue.PushFront(root)
	for queue.Len() > 0 {
		var current []int
		queueLen := queue.Len()
		for i := 0; i < queueLen; i++ {
			node := queue.Remove(queue.Back()).(*TreeNode)
			current = append(current, node.Val)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
		result = append(result, current)

	}
	return result

}

// 内存最小， 广度优先算法
func levelOrder44(root *TreeNode) [][]int {
	res := make([][]int, 0, 3)
	if root == nil {
		return res
	}
	//q := []*TreeNode{root}
	q := make([]*TreeNode, 1, 1)
	q[0] = root
	//q = append(q,root)
	for len(q) != 0 {
		ll := len(q)
		tmp := make([]int, 0, 3)
		for i := 0; i < ll; i++ {
			qi := q[0]
			//q = q[1:]
			q = append(q[:0], q[1:]...)
			tmp = append(tmp, qi.Val)
			if qi.Left != nil {
				q = append(q, qi.Left)
			}
			if qi.Right != nil {
				q = append(q, qi.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}
