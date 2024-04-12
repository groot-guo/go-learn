package main

import (
	"fmt"
	"time"
)

type Node struct {
	Val  int
	Next *Node
}

func reversal(root *Node) {
	if root == nil {
		return
	}
	p := root
	var prev *Node
	for p.Next != nil {
		tmp1 := p.Next
		p.Next = prev
		prev = p
		p = tmp1
	}
}

func main() {
	nowTime := time.Now()
	for i := 1; i < 100000; i++ {
		fmt.Println(i)
	}
	fmt.Println(time.Now().Unix() - nowTime.Unix())
	// res := dpSclie([]int{1, 1, 1, 2, 2, 2, 3, 3, 4})
	// fmt.Println(res)
}

func dpSclie(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	l := 0
	for i := 1; i < len(nums); i++ {
		if nums[l] == nums[i] {
			continue
		}
		nums[l+1] = nums[i]
		l++
		fmt.Println("---", l)
	}

	// existedMap := make(map[int]bool, 0)
	// for _, v := range nums {
	// 	if ok := existedMap[v]; ok {
	// 		continue
	// 	}
	// 	existedMap[v] = true
	// }

	// ans := make([]int, 0, len(nums))
	// for k, _ := range existedMap {
	// 	ans = append(ans, k)
	// }

	return nums[:l+1]
}

func getNum(n int) int {
	return 0
}
