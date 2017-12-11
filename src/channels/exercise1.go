package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	/*
		go func() {
			c <- 42
		}()
	*/

	c <- 42

	fmt.Println(<-c)
}

//https://play.golang.org/p/Mzn1Fkd0a6