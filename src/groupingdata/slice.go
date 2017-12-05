package main

import "fmt"

func main() {
	// COMPOSITE LITERAL
	/*
		x := []int{4, 5, 7, 8, 42}
		fmt.Println(x)
	*/
	// a SLICE allows you to group together VALUES of the same TYPE

	/*
		x := []int{4, 5, 7, 8, 42}
		fmt.Println(len(x))
		fmt.Println(x)
		fmt.Println(x[0])
		fmt.Println(x[1])
		fmt.Println(x[2])
		fmt.Println(x[3])
		fmt.Println(x[4])

		for i, v := range x {
			fmt.Println(i, v)
		}
	*/

	/*
		x := []int{4, 5, 7, 8, 42}
		fmt.Println(x[1])
		fmt.Println(x)
		fmt.Println(x[1:])
		fmt.Println(x[1:3])

		for i, v := range x {
			fmt.Println(i, v)
		}

		for i:= 0; i <= 4; i++ {
			fmt.Println(i, x[i])
		}
	*/

	/*
		x := []int{4, 5, 7, 8, 42}
		fmt.Println(x)
		x = append(x, 77, 88, 99, 1014)
		fmt.Println(x)

		y := []int{234, 456, 678, 987}
		x = append(x, y...)
		fmt.Println(x)
	*/

	/*
		x := []int{4, 5, 7, 8, 42}
		fmt.Println(x)
		x = append(x, 77, 88, 99, 1014)
		fmt.Println(x)

		y := []int{234, 456, 678, 987}
		x = append(x, y...)
		fmt.Println(x)

		x = append(x[:2], x[4:]...)
		fmt.Println(x)
	*/

	/*
		x := make([]int, 10, 12)
		fmt.Println(x)
		fmt.Println(len(x))
		fmt.Println(cap(x))
		x[0] = 42
		x[9] = 999

		fmt.Println(x)
		fmt.Println(len(x))
		fmt.Println(cap(x))

		x = append(x, 3423)

		fmt.Println(x)
		fmt.Println(len(x))
		fmt.Println(cap(x))

		x = append(x, 3424)

		fmt.Println(x)
		fmt.Println(len(x))
		fmt.Println(cap(x))

		x = append(x, 3425)

		fmt.Println(x)
		fmt.Println(len(x))
		fmt.Println(cap(x))
	*/

	/*
		jb := []string{"James", "Bond", "Chocolate", "martini"}
		fmt.Println(jb)

		mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
		fmt.Println(mp)

		xp := [][]string{jb, mp}
		fmt.Println(xp)
	*/

	/*
		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
		fmt.Println(x)

		y := append(x, 43, 43, 43, 43, 43, 43, 44) // new underlying array allocated

		fmt.Println(x)
		fmt.Println(y)
	*/

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	y := append(x[:2], x[5:]...) // the same underlying array stores the value of the new slice

	fmt.Println(x)
	fmt.Println(y)
}


//https://play.golang.org/p/O5Ofm5kRVG