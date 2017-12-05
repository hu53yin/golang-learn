package main

import (
	"fmt"
)

func main() {
	/*for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the remainder (aka modulus) is %v\n", i, i%4)
	}*/

	/*x := "James Bond"

	if x == "James Bond" {
		fmt.Println(x)
	}*/

	/*x := "Huseyin"

	if x == "James Bond" {
		fmt.Println(x)
	} else if x == "Huseyin" {
		fmt.Println(x, "Gurkan")
	} else {
		fmt.Println("Someone else")
	}*/

	/*switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}*/

	/*favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}*/

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

//https://play.golang.org/p/G3seu226Ok