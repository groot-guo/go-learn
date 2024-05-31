package main

/*
	十六进制 字符串输入 转十进制

	A   10
	AA

	01  1
	10  2
	11  3
	1011    2^1 * 2^0^2
*/

func fib(n int) int {
	p, q, r := 0, 0, 1

	if n < 3 {
		return r
	}

	for i := 1; i < n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func trap(height []int) int {
	leftMaxSlice := make([]int, len(height))
	left, right := 0, len(height)-1
	for i := 1; i < len(height); i++ {
		if height[left] > height[i] {
			leftMaxSlice[i] = height[left] - height[i]
		} else {
			left = i
		}
	}

	rightMaxSlice := make([]int, len(height))
	for i := right - 1; i >= 0; i-- {
		if height[right] > height[i] {
			rightMaxSlice[i] = height[right] - height[i]
		} else {
			right = i
		}
	}

	resultMAx := 0
	for l, r := 0, 0; l < len(leftMaxSlice) && r < len(rightMaxSlice); {

		resultMAx += min(leftMaxSlice[l], rightMaxSlice[r])
		l++
		r++
	}

	return resultMAx
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findAnagrams(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}

func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	//sort.Ints(nums)
	count, num := 0, k
	for i := range nums {
		for l := i; l < len(nums); l++ {
			num -= nums[l]
			if num == 0 {
				count++
			}
		}
	}

	return count
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {

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

	//var prev, node *ListNode
	//top := *head
	//node = head
	//prev = reversal(&top)
	//for prev != nil && node != nil {
	//	if prev.Val != node.Val {
	//		return false
	//	}
	//	prev = prev.Next
	//	node = node.Next
	//}
	//
	//return true
}

func reversal1(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

//func subarraySum(nums []int, k int) int {
//	count := 0
//	for start := 0; start < len(nums); start++ {
//		sum := 0
//		for end := start; end >= 0; end-- {
//			sum += nums[end]
//			if sum == k {
//				count++
//			}
//		}
//	}
//	return count
//}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	first, second := head, head.Next
	prev := &ListNode{Next: second}
	for second != nil {
		tmp := first
		second.Next = tmp
		first.Next = tmp.Next.Next

		if first == nil {
			break
		}
		first = tmp
		second = tmp.Next
	}

	return prev.Next
}

type Node1 struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	nodeList := make([]*Node, 0)
	prev, node := head, &Node{Val: head.Val}
	tmp := node
	for prev != nil {
		next := &Node{}
		tmp.Next = next
		prev = prev.Next
		next.Val = prev.Val
		nodeList = append(nodeList, tmp)
		tmp = next
	}

	nodeIndexMap := make(map[int]*Node)
	for _, v := range nodeList {
		nodeIndexMap[v.Val] = v
	}

	//temp, prev2 := node, head
	//for prev2 != nil {
	//	if prev2.Random != nil {
	//		key := nodeIndexMap[prev2.Random.Val]
	//		temp.Random = key
	//	}
	//	temp = temp.Next
	//	prev2 = prev2.Next
	//}

	return node
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}

	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}

	return dummyHead.Next
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(sort(head, mid), merge(mid, tail))
}
