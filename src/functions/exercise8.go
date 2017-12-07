package main

import (
	"fmt"
)

func main() {
	x := foo()
	x()
}

func foo() func() {
	return func() {
		fmt.Println("Hey foo func")
	}
}

//https://play.golang.org/p/fFKFBofDT8