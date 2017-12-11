package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cd\t%T\n", cs)

	// specific to general doesn't assign
	/*
		c = cr
		c = cs
	*/

	// general to specific assigns
	/*
		cr = c
		cs = c
	*/

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))
}

//https://play.golang.org/p/EH5wAfQ3pE