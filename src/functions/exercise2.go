package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}

func foo(i ...int) int {
	sum := 0
	for _, j := range i {
		sum += j
	}
	return sum
}

func bar(i []int) int {
	sum := 0
	for _, j := range i {
		sum += j
	}
	return sum
}

//https://play.golang.org/p/ex-h3IGjLq