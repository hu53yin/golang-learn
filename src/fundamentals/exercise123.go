package main

import (
	"fmt"
)

const (
	g = 42
	h int = 43
)

func main() {
	x := 42
	
	fmt.Printf("%d\t%b\t%#X\n", x, x, x)
	
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	
	fmt.Println(a, b, c, d, e, f)
	
	fmt.Println(g, h)
}

//https://play.golang.org/p/_vVcfvhWgv