package main

import (
	"fmt"
	"go-init-example/example_a"
	"go-init-example/example_b"
	"go-init-example/example_c"
	"go-init-example/example_c/example_d"
	"sync"
)

func main() {
	example_d.D = 4
	fmt.Println("d = ", example_d.D)

	fmt.Println("before a = ", example_a.A)
	example_a.A = 2
	fmt.Println("after a = ", example_a.A)
	fmt.Println("a1 = ", example_a.A1)

	example_b.B = 2
	fmt.Println("b = ", example_b.B)

	example_c.C = 3
	fmt.Println("c = ", example_c.C)

	wg := sync.WaitGroup{}
	s := sync.Once{}
	for i := 0; i < 3; i++ {
		go func() {
			s.Do(setA)
			fmt.Println("package a set A = ", A)
			wg.Done()
		}()
		wg.Add(1)
	}

	s.Do(setB)
	setA()

	wg.Wait()
}

var A int

func setA() {
	A++
	fmt.Println("set A++, A =  ", A)
}

func setB() {
	A++
	fmt.Println("set A++, A =  ", A)
}
