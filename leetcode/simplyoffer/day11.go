package simplyoffer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	18.删除链表的节点
	[4,5,1,9] val = 5
	搭建临时节点，判断下一个节点是否与删除元素相等，然后进行删除
*/

func deleteNode(head *ListNode, val int) *ListNode {
	tmp := &ListNode{Val: -1, Next: head}

	for ptr := tmp.Next; ptr.Next != nil; {
		if ptr.Next.Val == val {
			ptr.Next = ptr.Next.Next
		}
		ptr = ptr.Next
	}

	return tmp.Next
}

/*
22.链表中倒数第k个节点
1-2-3-4-5 和 k=2， 返回 4-5
思路： 1. 使用数组存放，2.链表反转,3.反转判断剩余链表
*/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	var reverseHead *ListNode

	for ptr := head; ptr != nil; {
		cur := ptr.Next
		ptr.Next = reverseHead
		reverseHead = ptr
		ptr = cur
	}

	var data *ListNode
	for k > 0 {
		tmp := reverseHead.Next
		cur := &ListNode{Val: reverseHead.Val}
		cur.Next = data
		reverseHead = tmp
		data = cur
		k--
	}
	return data
}

// 第3种 判断剩余数量, 1-2, 3-4-5,   1-2-3, 4-5
func getKthFromEnd2(head *ListNode, k int) *ListNode {
	left := head
	right := left
	for k > 0 {
		if right == nil {
			break
		}
		right = right.Next
		k--
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	return left
}
