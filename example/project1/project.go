package main

import (
	"fmt"
	"runtime"
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
	// nowTime := time.Now()
	// for i := 1; i < 100000; i++ {
	// 	fmt.Println(i)
	// }
	// fmt.Println(time.Now().Unix() - nowTime.Unix())
	res := dpSclie([]int{1, 1, 1, 2, 2, 2, 3, 3, 4})
	fmt.Println(res)
	ans := runtime.GOMAXPROCS(0)
	fmt.Println(ans)

	arr := make([]int, 0)
	//last := 0
	for i := 0; i < 2000; i++ {
		arr = append(arr, i)
		fmt.Println(len(arr), cap(arr))
	}

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
	asn := [2]int{}
	ss := make(map[[2]int]int)
	ss[asn] = 1
	fmt.Println(ss, asn)

	asn[0] = 1
	fmt.Println(ss, asn, ss[asn])
	// ss2 := make(map[map[string]string]string)
	// ss1 := &asn
	// ss3 := *ss1

	return nums[:l+1]
}

func getNum(n int) int {
	return n
}
