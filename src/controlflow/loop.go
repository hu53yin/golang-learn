package main

import "fmt"

func main() {
	// for init; condition; post {}
	/*for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}*/

	/*for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)

		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)
		}
	}*/

	/*x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}

	fmt.Println("done.")*/

	/*x := 1
	for {
		if x > 9 {
			break
		}

		fmt.Println(x)
		x++
	}
	fmt.Println("done.")*/

	/*x := 83 / 40
	y := 83 % 40
	fmt.Println(x, y)*/

	/*x := 1
	for {
		x++
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}

		fmt.Println(x)
	}
	fmt.Println("done.")*/
	
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
}

//https://play.golang.org/p/0mF2L7Ii-b