package main

import (
	"fmt"
	"math/rand"
)

/*
判断链表是否有环：给定一个链表头结点head，返回链表开始入环的第一个节点
*/

func main() {
	rand.Shuffle(10, func(i, j int) {

	})
	n := int64(10)
	rand.Seed(int64(n))
	fmt.Println(n)
}

type Node struct {
	Val  int
	Next *Node
}

func getHeadNode(node *Node) *Node {
	if node.Next == nil || node == nil {
		return nil
	}

	left, right := node, node.Next
	for left != right {
		if right == nil {
			break
		}
		left = left.Next
		right = right.Next
	}

	if right == nil {
		return nil
	}

	cur := left.Next
	nodeMap := make(map[*Node]struct{})
	nodeMap[left] = struct{}{}
	for cur != left {
		nodeMap[cur] = struct{}{}
		cur = cur.Next
	}

	ptr := node
	for {
		_, ok := nodeMap[ptr]
		if ok {
			break
		}
		ptr = ptr.Next
	}

	return ptr
}
