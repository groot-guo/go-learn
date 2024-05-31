package example_d

import (
	"fmt"
)

var D int

func init() {
	D = 1
	fmt.Println("package d set D = ", D)

	//example_a.A = 1
	//fmt.Println("package d set a = ", example_a.A)
}
