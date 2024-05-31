package main

import "fmt"

func fizzBuzz(n int) []string {
	res := make([]string, 0, n)
	for index := 0; index < n; index++ {
		sub := index + 1
		five := sub % 5
		three := sub % 3
		if five == 0 && three == 0 {
			res = append(res, "FizzBuzz")
		} else if five == 0 {
			res = append(res, "Buzz")
		} else if three == 0 {
			res = append(res, "Fizz")
		} else {
			res = append(res, fmt.Sprintf("%d", sub))
		}
	}
	return res
}
