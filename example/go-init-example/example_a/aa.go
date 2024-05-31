package example_a

import "fmt"

var A1 int

func init() {
	A1 = 11
	fmt.Println("package a set A1 = ", A1)
}
