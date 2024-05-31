package example_c

import (
	"fmt"
	"go-init-example/example_a"

	"go-init-example/example_c/example_d"
)

var C int

func init() {
	C = 1
	fmt.Println("package c set C = ", C)
	example_d.D = 10
	fmt.Println("package c set d = ", example_d.D)

	example_a.SetA()
	fmt.Println("package c set a = ", example_a.A)
}
