package main

import (
	"fmt"
)

var x, y int

func main() {
	x = 42
	y = 43
	fmt.Println(x, y)

	x, y = 44, 45
	fmt.Println(x, y)
}

//https://play.golang.org/p/o81FeqGBRM