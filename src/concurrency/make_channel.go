package main

import (
	"fmt"
)

func doSomething(x int) int {
	return x * 5
}

func main() {
	//make channel
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
}

//https://play.golang.org/p/8KwIkdGF5m