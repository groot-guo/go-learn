package simplyoffer

/*
25. 合并两个排序的链表
1-2-4, 1-3-4
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	left, right := l1, l2
	tmp := &ListNode{}
	ptr := tmp
	for {
		if left == nil {
			ptr.Next = right
			break
		}
		if right == nil {
			ptr.Next = left
			break
		}

		if left.Val <= right.Val {
			ptr.Next = left
			left = left.Next
		} else {
			ptr.Next = right
			right = right.Next
		}
		ptr = ptr.Next
	}

	return tmp.Next
}

/*
52.两个链表的第一个公共节点
listA = [4,1,8,4,5] listB = [5,0,1,8,4,5] skipA = 2, skipB = 3
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	left, right := headA, headB
	for left != right {
		// 这里只能使用 if else 判断
		if left == nil {
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
