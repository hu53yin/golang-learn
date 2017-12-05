package main

import (
	"fmt"
)

func main() {
	/*
		x := [5]int{1, 2, 3, 4, 5}

		for i, v := range x {
			fmt.Println(i, v)
		}

		fmt.Printf("%T\n", x)
	*/

	/*
		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

		for i, v := range x {
			fmt.Println(i, v)
		}

		fmt.Printf("%T\n", x)
	*/

	/*
		[42 43 44 45 46]
		[47 48 49 50 51]
		[44 45 46 47 48]
		[43 44 45 46 47]
	*/

	/*
		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

		fmt.Println(x[:5])
		fmt.Println(x[5:])
		fmt.Println(x[2:7])
		fmt.Println(x[1:6])
	*/

	/*
		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

		x = append(x, 52)
		fmt.Println(x)

		x = append(x, 53, 54, 55)
		fmt.Println(x)

		y := []int{56, 57, 58, 59, 60}

		x = append(x, y...)
		fmt.Println(x)
	*/

	/*
		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
		y := append(x[:3], x[6:]...)
		fmt.Println(y)
	*/

	states := make([]string, 50, 50)
	states = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println(len(states))
	fmt.Println(cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}
}

//https://play.golang.org/p/aFWTWlx2Th
