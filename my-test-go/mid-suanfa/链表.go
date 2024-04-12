package mid_suanfa

import "fmt"

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
	两数相加
*/
// 可以优化
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	data := &ListNode{}
	ptr := data
	temp := 0
	for l1 != nil && l2 != nil {
		cur := &ListNode{}
		val := l1.Val + l2.Val + temp
		if val < 9 {
			cur.Val = val
			temp = 0
		} else {
			cur.Val = val % 10
			temp = val / 10
		}
		l1 = l1.Next
		l2 = l2.Next
		ptr.Next = cur
		ptr = ptr.Next
	}

	for l1 != nil {
		cur := &ListNode{}
		val := l1.Val + temp
		if val < 9 {
			cur.Val = val
			temp = 0
		} else {
			cur.Val = val % 10
			temp = val / 10
		}
		ptr.Next = cur
		l1 = l1.Next
		ptr = ptr.Next
	}

	for l2 != nil {
		cur := &ListNode{}
		val := l2.Val + temp
		if val < 9 {
			cur.Val = val
			temp = 0
		} else {
			cur.Val = val % 10
			temp = val / 10
		}
		ptr.Next = cur
		l2 = l2.Next
		ptr = ptr.Next
	}
	if temp > 0 {
		ptr.Next = &ListNode{Val: temp}
	}

	return data
}

/*
	奇偶链表
*/

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	first, second := head, head.Next
	even := second
	for even != nil && even.Next != nil {
		first.Next = even.Next
		first = first.Next
		even.Next = first.Next
		even = even.Next
	}

	first.Next = second
	return head
}

/*
相交链表
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 双指针
	if headA == nil || headB == nil {
		return nil
	}
	left, right := headA, headB
	for left != right {
		fmt.Printf("l:%+v, r:%+v\n", left, right)
		if left == nil {
			// 因为出现说明没有相交
			left = headB
		} else {
			left = left.Next
		}
		if right == nil {
			right = headA
		} else {
			right = right.Next
		}
	}

	return left
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	existedMap := make(map[*ListNode]struct{})
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		existedMap[tmp] = struct{}{}
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if _, ok := existedMap[tmp]; ok {
			return tmp
		}
	}
	return nil
}
