package main

import "math"

type MinStack struct {
	topIndex int
	data     []int
	minVal   int
}

func Constructor22() MinStack {
	return MinStack{
		data:     make([]int, 0),
		topIndex: -1,
		minVal:   0,
	}
}

func (this *MinStack) Push(val int) {
	if this.topIndex+1 != len(this.data) {
		this.data[this.topIndex+1] = val
	} else {
		this.data = append(this.data, val)
	}
	this.topIndex++
}

func (this *MinStack) Pop() {
	this.topIndex--
}

func (this *MinStack) Top() int {
	return this.data[this.topIndex]
}

func (this *MinStack) GetMin() int {
	minVal := this.data[0]
	for index := 1; index <= this.topIndex; index++ {
		if minVal > this.data[index] {
			minVal = this.data[index]
		}
	}
	return minVal
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type MinStack22 struct {
	stack    []int
	minStack []int
}

func Constructor223() MinStack22 {
	return MinStack22{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack22) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack22) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack22) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack22) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}
