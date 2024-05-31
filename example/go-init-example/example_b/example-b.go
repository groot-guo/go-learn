package example_b

import "fmt"

var B int

func init() {
	B = 1
	fmt.Println("package b set B = ", B)
}
