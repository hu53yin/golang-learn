package main

import (
	"fmt"
)

func main() {

	func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("done")
}

//https://play.golang.org/p/DQX3xEIcRe