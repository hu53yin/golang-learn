package main

import (
	"fmt"
)

func main() {

	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		
		if len(xi) == 1 {
			return xi[0]
		}
		
		return xi[0] + xi[len(xi)-1]
	}

	x := foo(g, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)
}

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

//https://play.golang.org/p/GYZsHRq_xW