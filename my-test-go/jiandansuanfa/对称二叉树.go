package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helperTree(root.Left, root.Right)
}

func helperTree(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && helperTree(left.Left, right.Right) && helperTree(left.Right, right.Left)
}

// 递归
func isSymmetric22(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lnode := root.Left
	rnode := root.Right
	return bf(lnode, rnode)
}
func bf(lnode *TreeNode, rnode *TreeNode) bool {
	if lnode == nil && rnode == nil {
		return true
	}
	if rnode == nil || lnode == nil || rnode.Val != lnode.Val {
		return false
	}
	l := bf(lnode.Left, rnode.Right)
	r := bf(lnode.Right, rnode.Left)
	return r && l

}

// 迭代
func isSymmetric44(root *TreeNode) bool {
	u, v := root, root
	q := []*TreeNode{}
	q = append(q, u)
	q = append(q, v)
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true
}
