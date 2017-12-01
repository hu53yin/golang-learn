package main

import "fmt"

type hus int
var x hus
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 10
	fmt.Println(x)
	
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}

//https://play.golang.org/p/CiRmLsrrTH