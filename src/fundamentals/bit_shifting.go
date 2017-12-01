package main

import (
	"fmt"
)

const (
	_ = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)
	
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)
	
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}

//https://play.golang.org/p/HV3iispXvG