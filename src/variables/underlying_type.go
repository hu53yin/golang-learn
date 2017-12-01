package main

import "fmt"

type hus int

var x hus

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}

//https://play.golang.org/p/CFq10nDw0W