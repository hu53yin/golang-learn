package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	f := func() {
		fmt.Println("my first func expression")
	}

	f()

	g := func(x int) {
		fmt.Println("the year big brother started watching:", x)
	}

	g(1984)
}

//https://play.golang.org/p/ZMTWOG9QFw