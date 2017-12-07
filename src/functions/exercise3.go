package main

import (
	"fmt"
)

func main() {
	defer foo()
	fmt.Println("Hello, playground")
}

func foo() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}

//https://play.golang.org/p/XluEuUD0Nw