package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Custom error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "Hey wazz up!",
	}
	
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}

//https://play.golang.org/p/8JP8w8P2pN