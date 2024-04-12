package medium

import "leetcode/my-test/mid-suanfa"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
二叉树的中序遍历
*/
func inorderTraversal(root *mid_suanfa.TreeNode) []int {
	res := make([]int, 0)
	var helper func(node *mid_suanfa.TreeNode)
	helper = func(node *mid_suanfa.TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		res = append(res, node.Val)
		helper(node.Right)
	}
	helper(root)
	return res
}

func inorderTraversal2(root *mid_suanfa.TreeNode) []int {
	res := make([]int, 0)
	var stack []*mid_suanfa.TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}

	return res
}

/*
	二叉树的锯齿形层次遍历
*/

// bfs 广度优先遍历
func zigzagLevelOrder(root *mid_suanfa.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := []*mid_suanfa.TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		vals := []int{}
		q := queue
		queue = nil
		for _, node := range q {
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if level%2 == 1 {
			for i, n := 0, len(vals); i < n/2; i++ {
				vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
			}
		}
		res = append(res, vals)
	}

	return res
}

/*
	从前序与中序遍历序列构造二叉树
*/

// 递归
func buildTree(preorder []int, inorder []int) *mid_suanfa.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &mid_suanfa.TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if root.Val == inorder[i] {
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

// 迭代
func buildTree2(preorder []int, inorder []int) *mid_suanfa.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &mid_suanfa.TreeNode{Val: preorder[0]}
	stack := []*mid_suanfa.TreeNode{root}
	var inorderIndex int
	for i := 1; i < len(preorder); i++ {
		preoderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &mid_suanfa.TreeNode{Val: preoderVal}
			stack = append(stack, node.Left)
		} else {
			// 此处for 循环判断，取栈顶判断
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &mid_suanfa.TreeNode{Val: preoderVal}
			stack = append(stack, node.Right)
		}
	}
	return root
}

/*
填充每个节点的下一个右侧节点指针
*/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 层次遍历
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		tmp := queue
		queue = nil
		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

// 迭代
func connect2(root *Node) *Node {
	if root == nil {
		return root
	}

	// 每次循环从该层的最左侧节点开始
	for tmp := root; root.Left != nil; tmp = tmp.Left {
		// 通过 Next 遍历这一层节点，为下一层的节点更新 Next 指针
		for node := tmp; node != nil; node = node.Next {
			// 左节点指向右节点
			node.Left.Next = node.Right

			// 右节点指向下一个左节点
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}

	return root
}

func isSubStructure(A *mid_suanfa.TreeNode, B *mid_suanfa.TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return helper(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func helper(A, B *mid_suanfa.TreeNode) bool {
	if A == nil || A.Val != B.Val {
		return false
	}

	ans := true
	if B.Left != nil {
		ans = ans && helper(A.Left, B.Left)
	}
	if B.Right != nil {
		ans = ans && helper(A.Right, B.Right)
	}
	return ans
}
