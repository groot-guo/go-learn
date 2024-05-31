package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//时间复杂度：O(N)O(N)，其中 NN 是链表中的节点数。最坏情况下我们需要遍历每个节点一次。
//空间复杂度：O(N)O(N)，其中 NN 是链表中的节点数。主要为哈希表的开销，最坏情况下我们需要将每个节点插入到哈希表中一次。

func hasCycle(head *ListNode) bool {
	var data = make(map[*ListNode]bool)
	for head != nil {
		if head.Next == nil {
			return false
		}
		if data[head] {
			return true
		}
		data[head] = true
		head = head.Next
	}

	return false
}

// 运行最快， 快慢指针就是 龟兔赛跑相遇，或者 两人跑步相遇
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	// 快慢指针
	slow := head
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
