package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 使用双指针， 错误
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nodes := make([]*ListNode, 0)
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}

	left := 0
	right := len(nodes) - 1
	for left <= right {
		nodes[right].Next = nodes[left].Next
		left++
		right--
	}

	return dummy
}

// 赋值逻辑需要注意， 指针构建
func reverseList23(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	length := getLength(head)
	dummy := &ListNode{head.Val, nil}
	cur := dummy

	head = head.Next
	for index := 1; index < length; index++ {
		next := head.Next
		head.Next = cur
		cur = head // 此处注意
		head = next
	}

	return cur
}

// 迭代
func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
