package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.age)
}

func main() {
	p := person{
		"Huseyin",
		"Gurkan",
		35,
	}

	p.speak()
}

//https://play.golang.org/p/0b1hxyAIHp