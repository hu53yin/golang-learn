package main

import (
	"fmt"
)

func main() {
	var answer1, answer2, answer3 string

	fmt.Println("Name: ")
	_, err := fmt.Scan(&answer1)

	if err != nil {
		panic(err)
	}

	fmt.Println("Fav Food: ")
	_, err = fmt.Scan(&answer2)

	if err != nil {
		panic(err)
	}

	fmt.Println("Fav Sport: ")
	_, err = fmt.Scan(&answer3)

	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)
}

//https://play.golang.org/p/reC-8FnEdx