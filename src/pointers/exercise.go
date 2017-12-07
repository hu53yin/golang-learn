package main

import (
	"fmt"
)

type person struct {
	first string
}

func (p *person) changeMe(name string) {
	(*p).first = name
}

func main() {
	p := person{
		first: "Huseyin",
	}

	fmt.Println(p.first)

	p.changeMe("Berkay")

	fmt.Println(p.first)
}

//https://play.golang.org/p/bwQJiGFszw