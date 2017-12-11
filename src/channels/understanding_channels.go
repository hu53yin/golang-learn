package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)

	/*
		go func() {
			c <- 42
		}()
	*/

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
}

//https://play.golang.org/p/pb0gEtHDBH