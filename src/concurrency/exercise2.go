package main

import (
	"fmt"
)

type Person struct {
	name string
}

func (p *Person) speak() {
	fmt.Println("Hey")
}

type Human interface {
	speak()
}

func saySomething(h Human) {
	h.speak()
}

func main() {
	p := Person{
		"Huseyin",
	}

	saySomething(&p)
}

//https://play.golang.org/p/vhMUCrpnSc