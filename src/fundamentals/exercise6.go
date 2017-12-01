package main

import (
	"fmt"
)

const (
	a = 2017 + iota
	b = 2017 + iota
	c = 2017 + iota
	d = 2017 + iota
)

func main() {
	fmt.Println(a, b, c, d)
}

//https://play.golang.org/p/vV2_ND4GYj