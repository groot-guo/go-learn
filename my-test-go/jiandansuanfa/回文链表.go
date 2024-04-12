package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome2(head *ListNode) bool {
	var data []*ListNode
	for head != nil {
		data = append(data, head)
		head = head.Next
	}

	left := 0
	right := len(data) - 1
	for left < right {
		if data[left].Val != data[right].Val {
			return false
		}
		left++
		right--
	}
	return true
}

// 运行最快
func isPalindrome3(head *ListNode) bool {
	// 前半部分
	p1 := head
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 后半部分
	p2 := reverse23(slow.Next)
	for p1 != slow.Next && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func reverse23(node *ListNode) *ListNode {
	var pre *ListNode
	for node != nil {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return pre
}

// 内存最小
func isPalindrome34(head *ListNode) bool {
	var li []int
	for head != nil {
		li = append(li, head.Val)
		head = head.Next
	}
	if len(li) == 1 {
		return true
	}
	i := 0
	j := 0
	if len(li)%2 == 0 {
		i = len(li)/2 - 1
		j = len(li) / 2
	} else {
		j = len(li) / 2
		i = j
	}
	for i >= 0 && j < len(li) {
		if li[i] != li[j] {
			return false
		}
		i--
		j++
	}
	return true
}

// 递归
func isPalindrome13(head *ListNode) bool {
	frontPointer := head
	var recursivelyCheck func(*ListNode) bool
	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}
	return recursivelyCheck(head)
}

func reverseList55(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// o(n) 空间 o(1)
func isPalindrome55(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 找到前半部分链表的尾节点并反转后半部分链表
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList55(firstHalfEnd.Next)

	// 判断是否回文
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	firstHalfEnd.Next = reverseList55(secondHalfStart)
	return result
}
