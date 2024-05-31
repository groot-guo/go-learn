package main

import (
	"context"
	"fmt"
	"sync"
	"unsafe"
)

/*
go编写一个可以从 1 到 n 输出代表这个数字的字符串的程序，但是：

如果这个数字可以被 3 整除，输出 "fizz"。
如果这个数字可以被 5 整除，输出 "buzz"。
如果这个数字可以同时被 3 和 5 整除，输出 "fizzbuzz"。
例如，当 n = 15，输出： 1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz。

假设有这么一个类：

class FizzBuzz {
  public FizzBuzz(int n) { ... }               // constructor
  public void fizz(printFizz) { ... }          // only output "fizz"
  public void buzz(printBuzz) { ... }          // only output "buzz"
  public void fizzbuzz(printFizzBuzz) { ... }  // only output "fizzbuzz"
  public void number(printNumber) { ... }      // only output the numbers
}
请你实现一个有四个线程的多线程版  FizzBuzz， 同一个 FizzBuzz 实例会被如下四个线程使用：

线程A将调用 fizz() 来判断是否能被 3 整除，如果可以，则输出 fizz。
线程B将调用 buzz() 来判断是否能被 5 整除，如果可以，则输出 buzz。
线程C将调用 fizzbuzz() 来判断是否同时能被 3 和 5 整除，如果可以，则输出 fizzbuzz。
线程D将调用 number() 来实现输出既不能被 3 整除也不能被 5 整除的数字。
*/

type FizzBuzz struct {
	n         int
	mutx      *sync.WaitGroup
	curNumber int
	curNChan  chan int
}

func NewFizzBuzz(n int) *FizzBuzz {
	f := &FizzBuzz{n: n, mutx: &sync.WaitGroup{}, curNumber: 0, curNChan: make(chan int, 1)}
	f.curNChan <- f.curNumber + 1
	return f
}

func (f *FizzBuzz) fizz() {

	for {
		select {
		case f.curNumber = <-f.curNChan:
			if f.curNumber%3 == 0 && f.curNumber%5 != 0 {
				f.print("fizz")
			}
		}
		f.curNChan <- f.curNumber
		if f.curNumber > f.n {
			break
		}
	}
	f.mutx.Done()

	//if f.curNumber%3 == 0 && f.curNumber%5 != 0 {
	//	f.print("fizz")
	//}

}

func (f *FizzBuzz) buzz() {

	for {
		select {
		case f.curNumber = <-f.curNChan:
			if f.curNumber%5 == 0 && f.curNumber%3 != 0 {
				f.print("buzz")
			}
		}
		f.curNChan <- f.curNumber
		if f.curNumber > f.n {
			break
		}
	}
	f.mutx.Done()
	//if f.curNumber%5 == 0 && f.curNumber%3 != 0 {
	//	f.print("buzz")
	//}
	//f.mutx.Done()
}

func (f *FizzBuzz) fizzbuzz() {

	for {
		select {
		case f.curNumber = <-f.curNChan:
			if f.curNumber%3 == 0 && f.curNumber%5 == 0 {
				f.print("fizzbuzz")
			}
		}
		f.curNChan <- f.curNumber
		if f.curNumber > f.n {
			break
		}
	}
	f.mutx.Done()
	//if f.curNumber%3 == 0 && f.curNumber%5 == 0 {
	//	f.print("fizzbuzz")
	//}
	//f.mutx.Done()
}

func (f *FizzBuzz) number() {
	for {
		select {
		case f.curNumber = <-f.curNChan:
			if f.curNumber%3 != 0 && f.curNumber%5 != 0 {
				f.print(f.curNumber)
			}
		}
		//fmt.Println("number:---")
		f.curNChan <- f.curNumber
		if f.curNumber > f.n {
			break
		}
	}
	f.mutx.Done()
	//if f.curNumber%3 != 0 && f.curNumber%5 != 0 {
	//	f.print(f.curNumber)
	//}
	//f.mutx.Done()
}

func (f *FizzBuzz) print(res interface{}) {
	fmt.Println(res)
	f.curNumber++
}

func (f *FizzBuzz) Run() {
	f.mutx.Add(4)
	go f.fizz()
	go f.buzz()
	go f.fizzbuzz()
	go f.number()

	f.mutx.Wait()
	fmt.Println("------")
}

func main() {
	fizzBizz := NewFizzBuzz(15)

	fizzBizz.Run()

	res := [100]FizzBuzz{}
	res2 := [100]*FizzBuzz{}
	res3 := &res2
	res4 := res3[:]
	res5 := &res4
	fmt.Println("a:", unsafe.Sizeof(res),
		"b:", unsafe.Sizeof(res2),
		"c:", unsafe.Sizeof(res3),
		"d:", unsafe.Sizeof(res4),
		"e:", unsafe.Sizeof(res5))

	a := make([]int, 0)
	a = append(a, 1)
	b := append(a, 2)
	a = append(a, 3)
	c := append(b, 4)
	fmt.Println(a, b, c)

	m := map[int]int{}
	data := m[1]
	fmt.Println("data:", data)

	fmt.Println(2 >> 1)
	fmt.Println(2 << 1)
	fmt.Println(3 << 1)
	fmt.Println(4 >> 1)
	fmt.Println(1 & 1)
	fmt.Println(2 | 1)
	fmt.Println(2 & 1)
	fmt.Println(0 << 3)
	ctx := context.Background()
	fmt.Println(unsafe.Sizeof(ctx))

	context.WithValue(ctx, "1", "")

	type ss int
	s := new(ss)
	*s++
	fmt.Println(s, *s)
	fmt.Println(unsafe.Sizeof(ctx))

	cd := make(chan []int, 2)
	close(cd)
	d := <-cd
	d2 := <-cd
	fmt.Println(d, cap(d), len(d), d2)
	//cd <- []int{}

	//cs := make(chan<- int, 0)
	//cs <- 1
	//fmt.Println()

	sll := []int{10}
	set(sll)
	fmt.Println(sll, unsafe.Sizeof(sll))

	fmt.Println(struct{}{} == struct{}{})
	fmt.Println(map[struct{}]int{})
	fmt.Println(map[interface{}]int{})
	fmt.Println(map[[10]int]int{})

	fmt.Println("return:", bss())
	fmt.Println("return:", *(csss()))

	q := make(chan<- int, 2)
	q <- 1
	q <- 2
	select {
	case q <- 3:
		fmt.Println("ok")
	default:
		fmt.Println("wrong")
	}

}

func set(data []int) {
	data[0] = 1
	data = append(data, 1)
	data[0] = 1
	for _, r := range data {
		fmt.Println("before: ", r)
		r = 2
		fmt.Println("after: ", r)
	}
	fmt.Println(data, unsafe.Sizeof(data))
}

func bss() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i //或者直接写成return
}

func csss() *int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return &i
}

func maosort(data []int) {
	// 优化， 添加 标记， 避免无意义的循环
	flagSort := true

	for i := 0; i < len(data) && flagSort; i++ {
		flagSort = false
		for j := len(data) - 1; j >= i; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
				flagSort = true
			}
		}
	}
}

func jiandansort(data []int) {
	for i := 0; i < len(data); i++ {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[min] > data[j] {
				min = j
			}
		}
		if min != i {
			data[min], data[i] = data[i], data[min]
		}
	}
}

func insertSort(data []int) {
	j := 0
	for i := 2; i < len(data); i++ {
		if data[i] < data[i-1] {
			tmp := data[0]
			data[0] = data[i]
			for j = i - 1; data[j] > data[0]; j-- {
				data[j+1] = data[j]
			}
			data[j+1] = tmp
		}
	}
}

func buildHeap(data []int, i, n int) {
	temp := data[i]
	for j := 2 * i; j <= n; j *= 2 {
		if j < n && data[j] < data[j+1] {
			j++
		}
		if temp >= data[j] {
			break
		}
		data[i] = data[j]
		i = j
	}
	data[i] = temp
}

func sheelSort(data []int) {

	for i := len(data) / 2; i > 0; i-- {
		buildHeap(data, i, len(data))
	}

	for i := len(data) - 1; i >= 0; i-- {
		data[0], data[i] = data[i], data[0]
		buildHeap(data, 0, i-1)
	}

}
