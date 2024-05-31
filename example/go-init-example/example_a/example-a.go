package example_a

import (
	"fmt"
	"sync"
)

var A int

func init() {
	s := sync.Once{}
	s.Do(SetA)

	fmt.Println("package a set A = ", A)

	//example_d.D = 1111
	//fmt.Println("package a set d = ", example_d.D)

}

func SetA() {
	A++
	fmt.Println("set A++, A =  ", A)
}
