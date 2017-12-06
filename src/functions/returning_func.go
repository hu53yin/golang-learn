package main

import (
	"fmt"
)

func main() {
	x := bar()

	fmt.Printf("%T\n", x)

	fmt.Println(x())

	fmt.Println(bar()())
}

func bar() func() int {
	return func() int {
		return 451
	}
}

//https://play.golang.org/p/n1QYCDwMs_