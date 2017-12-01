package main

import "fmt"

const (
	a = 43
	b = 42.78
	c = "James Bond"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

//https://play.golang.org/p/NbY3Ud4IVV
