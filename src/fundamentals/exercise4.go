package main

import "fmt"

var a int

func main() {
	a = 24
	fmt.Printf("%d\t%b\t%#X\n", a, a, a)

	b := a << 1

	fmt.Printf("%d\t%b\t%#X\n", b, b, b)
}

//https://play.golang.org/p/J3j420X65y