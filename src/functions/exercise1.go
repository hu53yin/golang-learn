package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())

	x, y := bar()
	fmt.Println(x, y)

}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 1, "bar"
}

//https://play.golang.org/p/jy6clA5qH1