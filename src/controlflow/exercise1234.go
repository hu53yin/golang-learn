package main

import (
	"fmt"
)

func main() {
	/*for x := 1; x <= 10000; x++ {
		fmt.Println(x)
	}*/

	/*for x := 65; x <= 90; x++ {
		fmt.Println(x)
		for y := 0; y < 3; y++ {
			fmt.Printf("\t%#U\n", x)
		}
	}*/

	/*year := 1982

	for year <= 2017 {
		fmt.Println(year)
		year++
	}*/

	year := 1982

	for {
		if year <= 2017 {
			fmt.Println(year)
			year++
		} else {
			break
		}
	}
}

//https://play.golang.org/p/aZQzyFSy6m