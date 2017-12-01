package main

import (
	"fmt"
	"runtime"
)

var x byte = 128
var y rune = 128

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}

//https://play.golang.org/p/cT4CkSk-CZ