package main

import "fmt"

func main() {
	s := "Hello, playground"
	fmt.Println(s)
	s = "Hello, Hawaii"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}

//https://play.golang.org/p/S0K7G9Meqx