package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 2
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

//https://play.golang.org/p/zHbPiGz_xw